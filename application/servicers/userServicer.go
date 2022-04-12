package servicers

import (
	"fmt"

	dtos "github.com/garcia-paulo/go-gin/application/dtos/user"
	"github.com/garcia-paulo/go-gin/domain/models"
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

func (s *UserServicer) CreateUser(user models.User) (dtos.UserResponse, error) {
	if err := user.HashPassword(); err != nil {
		return dtos.UserResponse{}, err
	}
	s.userRepository.CreateUser(&user)
	if user.ID == 0 {
		return dtos.UserResponse{}, fmt.Errorf("error when saving to database")
	}

	return dtos.NewUserResponse(user), nil
}

func (s *UserServicer) AuthenticateUser(user dtos.UserRequest) (dtos.UserResponse, error) {
	foundUser := s.userRepository.FindUserByUsername(user.Username)
	err := foundUser.Authenticate(user.Password)
	return dtos.NewUserResponse(foundUser), err
}
