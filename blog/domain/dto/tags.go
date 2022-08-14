package dto

import "time"

type CreateTagsRequest struct {
	PostID int64 `json:"post_id" uri:"post_id" binding:"required"`
}

type CreateTagsResponse struct {
	ID     int64  `json:"createdId"`
	Name   string `json:"name"`
	PostID int64  `json:"post_id"`
}

type DeleteTagsRequest struct {
	PostID int64 `json:"post_id" uri:"post_id" binding:"required"`
	TagID  int64 `json:"tag_id" uri:"tag_id" binding:"required"`
}

type GetTagByIDRequest struct {
	PostID int64 `json:"post_id" uri:"post_id" binding:"required"`
	TagID  int64 `json:"tag_id" uri:"tag_id" binding:"required"`
}

type PostsTags struct {
	PostID int64 `sql:"type:int REFERENCES posts(id)" json:"post_id"`
	TagID  int64 `sql:"type:int REFERENCES tags(id)" json:"tag_id"`
}

//Tag Represents the fields from the Tags Database
type Tag struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	PostID    int64     `sql:"type:int REFERENCES posts(id)" json:"post_id"`
	Post      Post      `json:"post"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type TagCreate struct {
	Name string `json:"name"`
}

type UpdateTagsRequest struct {
	PostID int64 `json:"post_id" uri:"post_id" binding:"required"`
	TagID  int64 `json:"tag_id" uri:"tag_id" binding:"required"`
}

type UpdateTagsBodyRequest struct {
	Name string `json:"name"`
}
