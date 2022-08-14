package dto

import (
	"time"

	"github.com/lib/pq"
)

type CreatePostRequest struct {
	AuthorID int64 `json:"author_id" uri:"user_id" binding:"required"`
}

type CreatePostResponse struct {
	ID        int64          `json:"createdId"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	AuthorID  int64          `json:"author_id"`
	CreatedAt time.Time      `json:"created_at"`
	Tag       pq.StringArray `json:"tags"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type DeletePostRequest struct {
	PostID   int64 `json:"post_id" uri:"post_id" binding:"required"`
	AuthorID int64 `json:"author_id" uri:"user_id" binding:"required"`
}

type GetPosts struct {
	Offset  int `json:"offset" form:"from"`
	LastIdx int `json:"last_idx" form:"to"`
}

type GetPostByIDRequest struct {
	PostID   int64 `json:"post_id" uri:"post_id" binding:"required"`
	AuthorID int64 `json:"author_id" uri:"user_id" binding:"required"`
}

//Post Represents the fields from the Post Database
type Post struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorID  int64     `sql:"type:int REFERENCES users(id)" json:"author_id"`
	Tags      []Tag     `gorm:"many2many:posts_tags;"`
	TagsID    int64     `sql:"type:int REFERENCES tags(id)" json:"tags_id"`
	Comments  []Comment `gorm:"many2many:posts_comments"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PostCreate struct {
	ID        int64          `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Author    User           `json:"author"`
	AuthorID  int64          `json:"author_id"`
	TagsID    int64          `json:"tags_id" binding:"required"`
	Tags      pq.StringArray `json:"tags"`
	Comments  pq.StringArray `json:"comments"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type UpdatePostRequest struct {
	PostID   int64 `json:"post_id" uri:"post_id" binding:"required"`
	AuthorID int64 `json:"author_id" uri:"user_id" binding:"required"`
}

type UpdatePostBodyRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
