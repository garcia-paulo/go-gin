//go:build wireinject
// +build wireinject

package main

import (
	middleware "github.com/garcia-paulo/go-gin/application/middlewares/auth"
	"github.com/garcia-paulo/go-gin/application/servicers"
	"github.com/garcia-paulo/go-gin/application/token"
	"github.com/garcia-paulo/go-gin/infra/config"
	"github.com/garcia-paulo/go-gin/infra/database"
	"github.com/garcia-paulo/go-gin/infra/repositories"
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/garcia-paulo/go-gin/presentation/routes"
	"github.com/google/wire"
)

func InitializeRoutes() *routes.Server {
	panic(wire.Build(
		config.NewConfig,
		database.NewDatabase,
		repositories.NewStudentRepository,
		repositories.NewUserRepository,
		middleware.NewAuthMiddleware,
		servicers.NewStudentServicer,
		servicers.NewUserServicer,
		controllers.NewStudentController,
		controllers.NewUserController,
		routes.NewStudentRoutes,
		routes.NewUserRoutes,
		routes.NewServer,
		token.NewTokenMaker,
	))
}
