package servicers

import (
	"fmt"
	"time"

	dtos "github.com/garcia-paulo/go-gin/application/dtos/user"
	"github.com/garcia-paulo/go-gin/application/token"
	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/repositories"
)

type UserServicer struct {
	userRepository *repositories.UserRepository
	pasetoMaker    *token.PasetoMaker
}

func NewUserServicer(userRepository *repositories.UserRepository, pasetoMaker *token.PasetoMaker) *UserServicer {
	return &UserServicer{
		userRepository: userRepository,
		pasetoMaker:    pasetoMaker,
	}
}

func (s *UserServicer) CreateUser(user models.User) (*dtos.UserResponse, error) {
	if err := user.HashPassword(); err != nil {
		return nil, err
	}
	s.userRepository.CreateUser(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("error when saving to database")
	}

	duration, _ := time.ParseDuration("42h")

	token, err := s.pasetoMaker.CreateToken(user.Username, duration)
	if err != nil {
		return nil, err
	}

	return dtos.NewUserResponse(user, token), nil
}

func (s *UserServicer) AuthenticateUser(user dtos.UserRequest) (*dtos.UserResponse, error) {
	foundUser := s.userRepository.FindUserByUsername(user.Username)
	err := foundUser.Authenticate(user.Password)
	if err != nil {
		return nil, err
	}

	duration, _ := time.ParseDuration("42h")

	token, err := s.pasetoMaker.CreateToken(foundUser.Username, duration)
	if err != nil {
		return nil, err
	}

	return dtos.NewUserResponse(foundUser, token), err
}
