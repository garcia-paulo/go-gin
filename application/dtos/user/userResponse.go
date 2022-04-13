package dtos_user

import (
	"time"

	"github.com/garcia-paulo/go-gin/domain/models"
	"gorm.io/gorm"
)

type UserResponse struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Username  string         `json:"username"`
	Token     string         `json:"token"`
}

func NewUserResponse(user models.User, token string) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Username:  user.Username,
		Token:     token,
	}
}
