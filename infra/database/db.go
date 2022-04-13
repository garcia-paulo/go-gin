package database

import (
	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *config.Config) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(config.DBSource))
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.Student{})
	DB.AutoMigrate(&models.User{})

	return DB
}
