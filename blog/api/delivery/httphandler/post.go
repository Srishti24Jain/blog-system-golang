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

type postHandler struct {
	postUsecase interfaces.PostUsecase
}

func NewPostHandler(e *gin.Engine, p interfaces.PostUsecase) {
	handler := postHandler{postUsecase: p}
	e.GET("api/user/:user_id/post/:post_id", handler.GetPostByIdHandler)
	e.GET("api/posts", handler.GetPostsHandler)
	e.POST("api/user/:user_id/create-post", handler.CreatePostHandler)
	e.PUT("api/user/:user_id/post/:post_id", handler.UpdatePostHandler)
	e.DELETE("api/user/:user_id/post/:post_id", handler.DeletePostHandler)
}

func (s *postHandler) GetPostByIdHandler(ctx *gin.Context) {
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

	req := new(dto.GetPostByIDRequest)

	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	post, err := s.postUsecase.GetPostById(ctx, req.PostID, req.AuthorID)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Data: post,
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

func (s *postHandler) GetPostsHandler(ctx *gin.Context) {
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

	req := new(dto.GetPosts)
	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}
	if req.LastIdx == 0 {
		req.LastIdx = 100
	}
	offset := req.Offset
	limit := req.LastIdx - offset
	if limit <= 0 {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "'to' is less than 'from'",
		}
		return
	} else if limit > 100 {
		limit = 100
	}

	studies, err := s.postUsecase.GetAllPosts(ctx, limit, req.Offset)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Data: studies,
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

func (s *postHandler) CreatePostHandler(ctx *gin.Context) {
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
	req := new(dto.CreatePostRequest)
	if err := ctx.ShouldBindUri(&req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	reqBody := new(dto.PostCreate)
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	resp, err := s.postUsecase.CreatePost(ctx, req.AuthorID, reqBody)
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

func (s *postHandler) UpdatePostHandler(ctx *gin.Context) {
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

	req := new(dto.UpdatePostRequest)
	if err := ctx.BindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	reqBody := new(dto.UpdatePostBodyRequest)
	if err := ctx.Bind(&reqBody); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	resp, err := s.postUsecase.UpdatePost(ctx, req.PostID, req.AuthorID, reqBody)
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

func (s *postHandler) DeletePostHandler(ctx *gin.Context) {
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
	req := new(dto.DeletePostRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}
	err := s.postUsecase.DeletePost(ctx, req.PostID, req.AuthorID)
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
