package dto

import "time"

type CreateUserResponse struct {
	ID        int64     `json:"createdId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserRequest struct {
	UserID int64 `json:"user_id" uri:"user_id" binding:"required"`
}

type GetUsers struct {
	Offset  int `json:"offset" form:"from"`
	LastIdx int `json:"last_idx" form:"to"`
}

type GetUserByIDRequest struct {
	UserID int64 `json:"user_id" uri:"user_id" binding:"required"`
}

//User Represents the fields from the User Database
type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UpdateUserBodyRequest struct {
	Name string `json:"name"`
}

type UpdateUserRequest struct {
	UserID int64 `json:"user_id" uri:"user_id" binding:"required"`
}
