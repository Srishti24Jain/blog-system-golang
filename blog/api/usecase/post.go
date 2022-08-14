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

type postUsecase struct {
	db *gorm.DB
}

func NewPostUsecase(db *gorm.DB) interfaces.PostUsecase {
	return &postUsecase{
		db: db,
	}
}

func (uc *postUsecase) GetPostById(ctx *gin.Context, postID, authorID int64) (*dto.Post, error) {
	var post dto.Post
	res := uc.db.Debug().Find(&post, dto.Post{ID: postID})
	if res.RecordNotFound() {
		return nil, errors.New("post does not exist")
	}

	if post.ID != 0 {
		err := uc.db.Model(&dto.User{}).Where("id = ?", authorID).Take(&post.Author).Error
		if err != nil {
			return &dto.Post{}, err
		}
	}
	err := uc.db.Debug().Model(&dto.Tag{}).Where("id=?", post.TagsID).Take(&post.Tags).Error
	if err != nil {
		return &dto.Post{}, err
	}

	return &post, nil
}

func (uc *postUsecase) GetAllPosts(ctx *gin.Context, limit int, offset int) (*[]dto.Post, error) {
	posts := []dto.Post{}
	err := uc.db.Model(&dto.Post{}).Limit(limit).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	if len(posts) > 0 {
		for i, _ := range posts {
			err := uc.db.Model(&dto.Post{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]dto.Post{}, err
			}
			err = uc.db.Debug().Model(&dto.Tag{}).Where("id=?", posts[i].TagsID).Take(&posts[i].Tags).Error
			if err != nil {
				return &[]dto.Post{}, err
			}
		}
	}

	return &posts, nil
}

func (uc *postUsecase) CreatePost(ctx *gin.Context, authorID int64, request *dto.PostCreate) (dto.CreatePostResponse, error) {
	tx := uc.db.Begin()
	if tx.Error != nil {
		return dto.CreatePostResponse{}, tx.Error
	}

	request.AuthorID = authorID
	post := &dto.Post{
		Title:     request.Title,
		Content:   request.Content,
		AuthorID:  request.AuthorID,
		TagsID:    request.TagsID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := uc.db.Debug().Create(&post).Error
	if err != nil {
		return dto.CreatePostResponse{}, err
	}

	res := tx.Commit()
	if res.Error != nil {
		return dto.CreatePostResponse{}, res.Error
	}

	if post.ID != 0 {
		err := uc.db.Debug().Model(&dto.Tag{}).Where("id = ?", request.TagsID).Take(&post.Tags).Error
		if err != nil {
			return dto.CreatePostResponse{}, err
		}
	}

	var tagsName []string
	for _, v := range post.Tags {
		tagsName = append(tagsName, v.Name)
	}

	return dto.CreatePostResponse{
		ID:        post.ID,
		Title:     request.Title,
		Content:   request.Content,
		Tag:       tagsName,
		AuthorID:  request.AuthorID,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}, nil
}

func (uc *postUsecase) UpdatePost(ctx *gin.Context, postID, authorID int64, request *dto.UpdatePostBodyRequest) (*dto.Post, error) {

	var resp dto.Post
	author := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if len(request.Title) != 0 {
		author["title"] = request.Title
	}

	res := uc.db.Model(&resp).Where("id=?", postID).Take(&resp).Updates(author)
	if res.Error != nil {
		return &dto.Post{}, res.Error
	}

	if resp.ID != 0 {
		err := uc.db.Debug().Model(&dto.User{}).Where("id = ?", resp.AuthorID).Take(&resp.Author).Error
		if err != nil {
			return &dto.Post{}, err
		}
	}

	return &resp, nil
}

func (uc *postUsecase) DeletePost(ctx *gin.Context, postID, authorID int64) error {
	res := uc.db.Model(&dto.Post{}).Where("id = ? and author_id=?", postID, authorID).Take(&dto.Post{}).Delete(&dto.Post{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func AddTag(db *gorm.DB, post *dto.Post, tag *dto.Tag) error {

	fmt.Println("t----", tag)
	res := db.Model(&post).Debug().Association("Tags").Append(tag)
	return res.Error
}

func CreateTag(db *gorm.DB, tagName string, postId int64) (*dto.Tag, error) {
	var tag dto.Tag
	res := db.FirstOrCreate(&tag, &dto.Tag{Name: tagName})
	if res.Error != nil {
		return nil, res.Error
	}

	if tag.ID != 0 {
		err := db.Debug().Model(&dto.Post{}).Where("id = ?", postId).Take(&tag.Post).Error
		if err != nil {
			return &dto.Tag{}, err
		}
	}
	return &tag, nil
}
