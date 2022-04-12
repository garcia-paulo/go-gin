package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username" validate:"min=4,max=16" gorm:"uniqueIndex"`
	HashedPassword string `json:"password" validate:"min=6,max=18"`
}

func (u *User) HashPassword() error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(u.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(pwd)
	return nil
}

func (u *User) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
}

func (u *User) Validate() error {
	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}
