package interfaces

import (
	"github.com/gin-gonic/gin"

	"blog/domain/dto"
)

type CommentsUsecase interface {
	GetCommentById(ctx *gin.Context, CommentID, postID int64) (*dto.Comment, error)
	CreateComment(ctx *gin.Context, CommentID int64, request *dto.Comment) (dto.CreateCommentsResponse, error)
	UpdateComments(ctx *gin.Context, CommentID, postID int64, requestBody *dto.UpdateCommentsBodyRequest) (*dto.Comment, error)
	DeleteComments(ctx *gin.Context, CommentID, postID int64) error
}
