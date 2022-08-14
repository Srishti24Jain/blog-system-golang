package interfaces

import (
	"github.com/gin-gonic/gin"

	"blog/domain/dto"
)

type PostUsecase interface {
	GetPostById(ctx *gin.Context, postID, authorID int64) (*dto.Post, error)
	GetAllPosts(ctx *gin.Context, limit int, offset int) (*[]dto.Post, error)
	CreatePost(ctx *gin.Context, authorID int64, request *dto.PostCreate) (dto.CreatePostResponse, error)
	UpdatePost(ctx *gin.Context, postID, authorID int64, requestBody *dto.UpdatePostBodyRequest) (*dto.Post, error)
	DeletePost(ctx *gin.Context, postID, authorID int64) error
}
