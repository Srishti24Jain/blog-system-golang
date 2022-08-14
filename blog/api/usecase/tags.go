package usecase

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"blog/domain/dto"
	"blog/domain/interfaces"
)

type tagsUsecase struct {
	db *gorm.DB
}

func NewTagsUsecase(db *gorm.DB) interfaces.TagsUsecase {
	return &tagsUsecase{
		db: db,
	}
}

func (uc *tagsUsecase) GetTagById(ctx *gin.Context, tagsID, postID int64) (*dto.Tag, error) {
	var tag dto.Tag
	res := uc.db.Debug().Find(&tag, dto.Tag{ID: tagsID})
	if res.RecordNotFound() {
		return nil, errors.New("post does not exist")
	}

	if tag.ID != 0 {
		err := uc.db.Model(&dto.Post{}).Where("id = ?", postID).Take(&tag.Post).Error
		if err != nil {
			return &dto.Tag{}, err
		}
	}
	fmt.Println("tags---", tag.Post.Tags)
	return &tag, nil
}

func (uc *tagsUsecase) CreateTag(ctx *gin.Context, postID int64, request *dto.Tag) (dto.CreateTagsResponse, error) {
	request.PostID = postID
	tag := &dto.Tag{
		Name:   request.Name,
		PostID: request.PostID,
	}

	err := uc.db.Debug().Create(&tag).Error
	if err != nil {
		return dto.CreateTagsResponse{}, err
	}

	return dto.CreateTagsResponse{
		ID:     tag.ID,
		Name:   request.Name,
		PostID: request.PostID,
	}, nil
}

func (uc *tagsUsecase) UpdateTags(ctx *gin.Context, tagID, postID int64, request *dto.UpdateTagsBodyRequest) (*dto.Tag, error) {
	var resp dto.Tag
	tag := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if len(request.Name) != 0 {
		tag["name"] = request.Name
	}

	res := uc.db.Model(&resp).Where("id=?", tagID).Take(&resp).Updates(tag)
	if res.Error != nil {
		return &dto.Tag{}, res.Error
	}

	if resp.ID != 0 {
		err := uc.db.Debug().Model(&dto.Post{}).Where("id = ?", postID).Take(&resp.Post).Error
		if err != nil {
			return &dto.Tag{}, err
		}
	}

	return &resp, nil
}

func (uc *tagsUsecase) DeleteTags(ctx *gin.Context, tagsID, PostID int64) error {
	res := uc.db.Model(&dto.Tag{}).Where("id = ? and post_id=?", tagsID, PostID).Take(&dto.Tag{}).Delete(&dto.Tag{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
