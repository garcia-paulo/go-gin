package database

import (
	"github.com/garcia-paulo/go-gin/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConnect() {
	connectionString := "host=localhost user=postgres password=postgres dbname=go-gin port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.Student{})
}
