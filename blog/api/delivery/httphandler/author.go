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

type userHandler struct {
	userUsecase interfaces.UserUsecase
}

func NewUserHandler(e *gin.Engine, a interfaces.UserUsecase) {
	handler := userHandler{userUsecase: a}
	e.GET("api/user/:user_id", handler.GetUserByIdHandler)
	e.GET("api/users", handler.GetUsersHandler)
	e.POST("api/create-user", handler.CreateUserHandler)
	e.PUT("api/user/:user_id", handler.UpdateUserHandler)
	e.DELETE("api/user/:user_id", handler.DeleteUserHandler)
}

func (s *userHandler) GetUserByIdHandler(ctx *gin.Context) {
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

	req := new(dto.GetUserByIDRequest)

	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	user, err := s.userUsecase.GetUserById(ctx, req.UserID)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Data: user,
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

func (s *userHandler) GetUsersHandler(ctx *gin.Context) {
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

	req := new(dto.GetUsers)
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

	users, err := s.userUsecase.GetAllUsers(ctx, limit, req.Offset)
	if err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusInternalServerError),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}
		return
	}

	data, err := json.Marshal(httputil.StandardEnvelope{
		Data: users,
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

func (s *userHandler) CreateUserHandler(ctx *gin.Context) {
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

	req := new(dto.User)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	resp, err := s.userUsecase.CreateUser(ctx, req)
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

func (s *userHandler) UpdateUserHandler(ctx *gin.Context) {
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

	req := new(dto.UpdateUserRequest)
	if err := ctx.BindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	reqBody := new(dto.UpdateUserBodyRequest)
	if err := ctx.Bind(&reqBody); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}

	resp, err := s.userUsecase.UpdateUser(ctx, req.UserID, reqBody)
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

func (s *userHandler) DeleteUserHandler(ctx *gin.Context) {
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
	req := new(dto.DeleteUserRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		httpError = &httputil.StandardError{
			Code:   strconv.Itoa(http.StatusBadRequest),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}
		return
	}
	err := s.userUsecase.DeleteUser(ctx, req.UserID)
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
