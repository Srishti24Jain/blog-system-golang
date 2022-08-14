package httphandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"blog/domain/dto"
	"blog/domain/interfaces"
	"blog/utils/httputil"
)

type tagsHandler struct {
	tagsUsecase interfaces.TagsUsecase
}

func NewTagsHandler(e *gin.Engine, a interfaces.TagsUsecase) {
	handler := tagsHandler{tagsUsecase: a}
	e.GET("api/post/:post_id/tags/:tag_id", handler.GetTagByIdHandler)
	e.POST("api/post/:post_id/create-tag", handler.CreateTagsHandler)
	e.PUT("api/post/:post_id/tags/:tag_id", handler.UpdateTagsHandler)
	e.DELETE("api/post/:post_id/tags/:tag_id", handler.DeleteTagsHandler)
}

func (s *tagsHandler) GetTagByIdHandler(ctx *gin.Context) {
	var (
		startTime = time.Now()
		httpError *httputil.StandardError
	)
	defer func() {
		if httpError != nil {
			errCode, _ := strconv.Atoi(httpError.Code)
			httputil.WriteErrorResponse(ctx.Writer, errCode, []httputil.StandardError{*httpError})
		}
	}()

	req := new(dto.GetTagByIDRequest)

	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	tag, err := s.tagsUsecase.GetTagById(ctx, req.TagID, req.PostID)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Data: tag,
		Status: &httputil.StandardStatus{
			Message:   http.StatusText(http.StatusOK),
			ErrorCode: 0,
		},
		Header: &httputil.StandardHeader{
			TotalData:   1,
			ProcessTime: time.Since(startTime).Seconds(),
		},
	})
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}
	_, _ = httputil.WriteJSONResponse(ctx.Writer, data, http.StatusOK)
	return
}

func (s *tagsHandler) CreateTagsHandler(ctx *gin.Context) {
	var (
		startTime = time.Now()
		httpError *httputil.StandardError
	)
	defer func() {
		if httpError != nil {
			errCode, _ := strconv.Atoi(httpError.Code)
			httputil.WriteErrorResponse(ctx.Writer, errCode, []httputil.StandardError{*httpError})
		}
	}()
	req := new(dto.CreateTagsRequest)
	if err := ctx.ShouldBindUri(&req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	reqBody := new(dto.Tag)
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	resp, err := s.tagsUsecase.CreateTag(ctx, req.PostID, reqBody)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Data: resp,
		Status: &httputil.StandardStatus{
			Message:   http.StatusText(http.StatusOK),
			ErrorCode: 0,
		},
		Header: &httputil.StandardHeader{
			TotalData:   1,
			ProcessTime: time.Since(startTime).Seconds(),
		},
	})
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}
	_, _ = httputil.WriteJSONResponse(ctx.Writer, data, http.StatusCreated)
	return
}

func (s *tagsHandler) UpdateTagsHandler(ctx *gin.Context) {
	var (
		startTime = time.Now()
		httpError *httputil.StandardError
	)
	defer func() {
		if httpError != nil {
			errCode, _ := strconv.Atoi(httpError.Code)
			httputil.WriteErrorResponse(ctx.Writer, errCode, []httputil.StandardError{*httpError})
		}
	}()

	req := new(dto.UpdateTagsRequest)
	if err := ctx.BindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	reqBody := new(dto.UpdateTagsBodyRequest)
	if err := ctx.Bind(&reqBody); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	resp, err := s.tagsUsecase.UpdateTags(ctx, req.TagID, req.PostID, reqBody)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Data: resp,
		Status: &httputil.StandardStatus{
			Message:   http.StatusText(http.StatusOK),
			ErrorCode: 0,
		},
		Header: &httputil.StandardHeader{
			TotalData:   0,
			ProcessTime: time.Since(startTime).Seconds(),
		},
	})
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}
	_, _ = httputil.WriteJSONResponse(ctx.Writer, data, http.StatusOK)
	return
}

func (s *tagsHandler) DeleteTagsHandler(ctx *gin.Context) {
	var (
		startTime = time.Now()
		httpError *httputil.StandardError
	)
	defer func() {
		if httpError != nil {
			errCode, _ := strconv.Atoi(httpError.Code)
			httputil.WriteErrorResponse(ctx.Writer, errCode, []httputil.StandardError{*httpError})
		}
	}()
	req := new(dto.DeleteTagsRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}
	err := s.tagsUsecase.DeleteTags(ctx, req.TagID, req.PostID)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Status: &httputil.StandardStatus{
			Message:   http.StatusText(http.StatusOK),
			ErrorCode: 0,
		},
		Header: &httputil.StandardHeader{
			TotalData:   1,
			ProcessTime: time.Since(startTime).Seconds(),
		},
	})
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}
	_, _ = httputil.WriteJSONResponse(ctx.Writer, data, http.StatusNoContent)
	return
}
