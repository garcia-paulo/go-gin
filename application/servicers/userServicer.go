package servicers

import (
	"fmt"

	input_user "github.com/garcia-paulo/go-gin/application/dtos/user/input"
	output_user "github.com/garcia-paulo/go-gin/application/dtos/user/output"
	"github.com/garcia-paulo/go-gin/application/token"
	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/config"
	"github.com/garcia-paulo/go-gin/infra/repositories"
)

type UserServicer struct {
	userRepository *repositories.UserRepository
	tokenMaker     *token.TokenMaker
	config         *config.Config
}

func NewUserServicer(userRepository *repositories.UserRepository, tokenMaker *token.TokenMaker, config *config.Config) *UserServicer {
	return &UserServicer{
		userRepository: userRepository,
		tokenMaker:     tokenMaker,
		config:         config,
	}
}

func (s *UserServicer) CreateUser(data input_user.UserRequest) (*output_user.UserResponse, error) {
	user := models.NewUser(data)
	if err := user.HashPassword(); err != nil {
		return nil, err
	}
	s.userRepository.CreateUser(user)
	if user.ID == 0 {
		return nil, fmt.Errorf("error when saving to database")
	}

	token, err := s.tokenMaker.CreateToken(user.Username, s.config.TokenDuration)
	if err != nil {
		return nil, err
	}

	return output_user.NewUserResponse(user, token), nil
}

func (s *UserServicer) AuthenticateUser(user input_user.UserRequest) (*output_user.UserResponse, error) {
	foundUser := s.userRepository.FindUserByUsername(user.Username)
	err := foundUser.Authenticate(user.Password)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenMaker.CreateToken(foundUser.Username, s.config.TokenDuration)
	if err != nil {
		return nil, err
	}

	return output_user.NewUserResponse(&foundUser, token), err
}
