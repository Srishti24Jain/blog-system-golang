package usecase

import (
	"time"

	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"blog/domain/dto"
	"blog/domain/interfaces"
)

type userUsecase struct {
	db *gorm.DB
}

func NewUserUsecase(db *gorm.DB) interfaces.UserUsecase {
	return &userUsecase{
		db: db,
	}
}

func (uc *userUsecase) GetUserById(ctx *gin.Context, userID int64) (*dto.User, error) {
	var user dto.User
	res := uc.db.Debug().Find(&user, &dto.User{ID: userID})
	if res.RecordNotFound() {
		return nil, errors.New("user does not exist")
	}

	return &user, nil
}

func (uc *userUsecase) GetAllUsers(ctx *gin.Context, limit int, offset int) (*[]dto.User, error) {
	users := []dto.User{}
	err := uc.db.Model(&dto.User{}).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (uc *userUsecase) CreateUser(ctx *gin.Context, request *dto.User) (dto.CreateUserResponse, error) {
	err := uc.db.Debug().Create(&request).Error
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}, nil
}

func (uc *userUsecase) UpdateUser(ctx *gin.Context, authorID int64, request *dto.UpdateUserBodyRequest) (*dto.User, error) {
	user := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if len(request.Name) != 0 {
		user["name"] = request.Name
	}
	res := uc.db.Model(&dto.User{}).Where("id=?", authorID).Take(&dto.User{}).UpdateColumns(user)
	if res.Error != nil {
		return &dto.User{}, res.Error
	}

	var resp dto.User
	err := res.Model(&dto.User{}).Where("id", authorID).Take(&resp).Error
	if err != nil {
		return &dto.User{}, err
	}

	return &resp, nil
}

func (uc *userUsecase) DeleteUser(ctx *gin.Context, userID int64) error {
	res := uc.db.Model(&dto.User{}).Where("id = ?", userID).Take(&dto.User{}).Delete(&dto.User{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
