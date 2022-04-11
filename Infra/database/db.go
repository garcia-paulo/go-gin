package database

import (
	"github.com/garcia-paulo/go-gin/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(connectionString string) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.Student{})

	return DB
}

func NewDatabase() *gorm.DB {
	return ConnectDatabase("host=localhost user=postgres password=postgres dbname=go-gin port=5432 sslmode=disable")
}
