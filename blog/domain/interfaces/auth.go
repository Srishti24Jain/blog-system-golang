package interfaces

import (
	"github.com/gin-gonic/gin"

	"blog/domain/dto"
)

type UserUsecase interface {
	GetUserById(ctx *gin.Context, userID int64) (*dto.User, error)
	GetAllUsers(ctx *gin.Context, limit int, offset int) (*[]dto.User, error)
	CreateUser(ctx *gin.Context, request *dto.User) (dto.CreateUserResponse, error)
	UpdateUser(ctx *gin.Context, userID int64, requestBody *dto.UpdateUserBodyRequest) (*dto.User, error)
	DeleteUser(ctx *gin.Context, userID int64) error
}
