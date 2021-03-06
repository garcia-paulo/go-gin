package controllers

import (
	"net/http"

	input_user "github.com/garcia-paulo/go-gin/application/dtos/user/input"
	"github.com/garcia-paulo/go-gin/application/servicers"
	"github.com/garcia-paulo/go-gin/application/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userServicer *servicers.UserServicer
}

func NewUserController(userServicer *servicers.UserServicer) *UserController {
	return &UserController{
		userServicer: userServicer,
	}
}

func (c *UserController) CreateUser(context *gin.Context) {
	user := input_user.UserRequest{}

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := user.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	response, err := c.userServicer.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, response)
}

func (c *UserController) AuthenticateUser(context *gin.Context) {
	user := input_user.UserRequest{}

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	response, err := c.userServicer.AuthenticateUser(user)
	if err != nil {
		context.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, response)
}
