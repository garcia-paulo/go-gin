package routes

import (
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController *controllers.UserController
}

func NewUserRoutes(userController *controllers.UserController) *UserRoutes {
	return &UserRoutes{
		userController: userController,
	}
}

func (r *UserRoutes) HandleRequests(g *gin.Engine) {
	g.POST("/users", r.userController.CreateUser)
	g.POST("/users/auth", r.userController.AuthenticateUser)
}
