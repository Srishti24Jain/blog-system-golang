package dto

import "time"

//Comment Represents the fields from the comments Database
type Comment struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	PostID    int64     `sql:"type:int REFERENCES posts(id)" json:"post_id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Body      string    `gorm:"size:255;not null" json:"body"`
	Post      Post      `json:"post"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CreateCommentsRequest struct {
	PostID int64 `json:"post_id" uri:"post_id" binding:"required"`
}

type CreateCommentsResponse struct {
	ID     int64  `json:"createdId"`
	Name   string `json:"name"`
	Body   string `json:"body"`
	PostID int64  `json:"post_id"`
}

type DeleteCommentRequest struct {
	PostID    int64 `json:"post_id" uri:"post_id" binding:"required"`
	CommentID int64 `json:"comment_id" uri:"comment_id" binding:"required"`
}

type GetCommentByIDRequest struct {
	PostID    int64 `json:"post_id" uri:"post_id" binding:"required"`
	CommentID int64 `json:"comment_id" uri:"comment_id" binding:"required"`
}

type UpdateCommentsBodyRequest struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

type UpdateCommentsRequest struct {
	PostID    int64 `json:"post_id" uri:"post_id" binding:"required"`
	CommentID int64 `json:"comment_id" uri:"comment_id" binding:"required"`
}
