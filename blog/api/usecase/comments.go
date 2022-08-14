package usecase

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"blog/domain/dto"
	"blog/domain/interfaces"
)

type commentsUsecase struct {
	db *gorm.DB
}

func NewCommentsUsecase(db *gorm.DB) interfaces.CommentsUsecase {
	return &commentsUsecase{
		db: db,
	}
}

func (uc *commentsUsecase) GetCommentById(ctx *gin.Context, commentID, postID int64) (*dto.Comment, error) {
	var comment dto.Comment
	res := uc.db.Debug().Find(&comment, dto.Comment{ID: commentID})
	if res.RecordNotFound() {
		return nil, errors.New("error")
	}

	if comment.ID != 0 {
		err := uc.db.Model(&dto.Post{}).Where("id = ?", postID).Take(&comment.Post).Error
		if err != nil {
			return &dto.Comment{}, err
		}
	}

	return &comment, nil
}

func (uc *commentsUsecase) CreateComment(ctx *gin.Context, postID int64, request *dto.Comment) (dto.CreateCommentsResponse, error) {
	request.PostID = postID
	comment := &dto.Comment{
		Name:   request.Name,
		Body:   request.Body,
		PostID: request.PostID,
	}

	err := uc.db.Debug().Create(&comment).Error
	if err != nil {
		return dto.CreateCommentsResponse{}, err
	}

	return dto.CreateCommentsResponse{
		ID:     comment.ID,
		Name:   request.Name,
		Body:   request.Body,
		PostID: request.PostID,
	}, nil
}

func (uc *commentsUsecase) UpdateComments(ctx *gin.Context, commentID, postID int64, request *dto.UpdateCommentsBodyRequest) (*dto.Comment, error) {
	var resp dto.Comment
	comment := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if len(request.Name) != 0 {
		comment["name"] = request.Name
	}

	if len(request.Body) != 0 {
		comment["body"] = request.Body
	}

	res := uc.db.Model(&resp).Where("id=?", commentID).Take(&resp).Updates(comment)
	if res.Error != nil {
		return &dto.Comment{}, res.Error
	}

	if resp.ID != 0 {
		err := uc.db.Debug().Model(&dto.Post{}).Where("id = ?", postID).Take(&resp.Post).Error
		if err != nil {
			return &dto.Comment{}, err
		}
	}

	return &resp, nil
}

func (uc *commentsUsecase) DeleteComments(ctx *gin.Context, commentID, PostID int64) error {
	res := uc.db.Model(&dto.Comment{}).Where("id = ? and post_id=?", commentID, PostID).Take(&dto.Comment{}).Delete(&dto.Comment{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
