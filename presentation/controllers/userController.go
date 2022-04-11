package controllers

import "github.com/garcia-paulo/go-gin/application/servicers"

type UserController struct {
	userServicer *servicers.UserServicer
}

func NewUserController(userServicer *servicers.UserServicer) *UserController {
	return &UserController{
		userServicer: userServicer,
	}
}
