package models

import (
	input_user "github.com/garcia-paulo/go-gin/application/dtos/user/input"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string
	HashedPassword string
}

func NewUser(user input_user.UserRequest) *User {
	return &User{
		Username:       user.Username,
		HashedPassword: user.Password,
	}
}

func (u *User) HashPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(password)
	return nil
}

func (u *User) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
}
