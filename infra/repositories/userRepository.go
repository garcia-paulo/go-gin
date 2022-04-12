package repositories

import (
	"github.com/garcia-paulo/go-gin/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (r *UserRepository) CreateUser(user *models.User) {
	r.database.Create(user)
}

func (r *UserRepository) FindUserByUsername(username string) models.User {
	user := models.User{}
	r.database.Where(models.User{Username: username}).First(&user)
	return user
}
