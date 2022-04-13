package servicers

import (
	"fmt"

	dtos "github.com/garcia-paulo/go-gin/application/dtos/user"
	"github.com/garcia-paulo/go-gin/application/token"
	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/config"
	"github.com/garcia-paulo/go-gin/infra/repositories"
)

type UserServicer struct {
	userRepository *repositories.UserRepository
	pasetoMaker    *token.PasetoMaker
	config         *config.Config
}

func NewUserServicer(userRepository *repositories.UserRepository, pasetoMaker *token.PasetoMaker, config *config.Config) *UserServicer {
	return &UserServicer{
		userRepository: userRepository,
		pasetoMaker:    pasetoMaker,
		config:         config,
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

	token, err := s.pasetoMaker.CreateToken(user.Username, s.config.TokenDuration)
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

	token, err := s.pasetoMaker.CreateToken(foundUser.Username, s.config.TokenDuration)
	if err != nil {
		return nil, err
	}

	return dtos.NewUserResponse(foundUser, token), err
}
