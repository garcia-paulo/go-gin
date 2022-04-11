package servicers

import (
	"github.com/garcia-paulo/go-gin/infra/repositories"
)

type UserServicer struct {
	userRepository *repositories.UserRepository
}

func NewUserServicer(userRepository *repositories.UserRepository) *UserServicer {
	return &UserServicer{
		userRepository: userRepository,
	}
}
