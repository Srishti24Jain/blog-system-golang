package interfaces

import (
	"github.com/gin-gonic/gin"

	"blog/domain/dto"
)

type TagsUsecase interface {
	GetTagById(ctx *gin.Context, tagID, postID int64) (*dto.Tag, error)
	CreateTag(ctx *gin.Context, tagID int64, request *dto.Tag) (dto.CreateTagsResponse, error)
	UpdateTags(ctx *gin.Context, tagID, postID int64, requestBody *dto.UpdateTagsBodyRequest) (*dto.Tag, error)
	DeleteTags(ctx *gin.Context, tagID, postID int64) error
}
