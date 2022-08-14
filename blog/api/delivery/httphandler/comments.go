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

type commentsHandler struct {
	commentsUsecase interfaces.CommentsUsecase
}

func NewCommentsHandler(e *gin.Engine, a interfaces.CommentsUsecase) {
	handler := commentsHandler{commentsUsecase: a}
	e.GET("api/post/:post_id/comments/:comment_id", handler.GetCommentByIdHandler)
	e.POST("api/post/:post_id/add-comment", handler.CreateCommentsHandler)
	e.PUT("api/post/:post_id/comments/:comment_id", handler.UpdateCommentsHandler)
	e.DELETE("api/post/:post_id/comments/:comment_id", handler.DeleteCommentsHandler)
}

func (s *commentsHandler) GetCommentByIdHandler(ctx *gin.Context) {
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

	req := new(dto.GetCommentByIDRequest)

	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	comment, err := s.commentsUsecase.GetCommentById(ctx, req.CommentID, req.PostID)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Data: comment,
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

func (s *commentsHandler) CreateCommentsHandler(ctx *gin.Context) {
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
	req := new(dto.CreateCommentsRequest)
	if err := ctx.ShouldBindUri(&req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	reqBody := new(dto.Comment)
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	resp, err := s.commentsUsecase.CreateComment(ctx, req.PostID, reqBody)
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

func (s *commentsHandler) UpdateCommentsHandler(ctx *gin.Context) {
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

	req := new(dto.UpdateCommentsRequest)
	if err := ctx.BindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	reqBody := new(dto.UpdateCommentsBodyRequest)
	if err := ctx.Bind(&reqBody); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	resp, err := s.commentsUsecase.UpdateComments(ctx, req.CommentID, req.PostID, reqBody)
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

func (s *commentsHandler) DeleteCommentsHandler(ctx *gin.Context) {
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
	req := new(dto.DeleteCommentRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}
	err := s.commentsUsecase.DeleteComments(ctx, req.CommentID, req.PostID)
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
