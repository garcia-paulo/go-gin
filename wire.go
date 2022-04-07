//go:build wireinject
// +build wireinject

package main

import (
	"github.com/garcia-paulo/go-gin/Application/servicers"
	"github.com/garcia-paulo/go-gin/Infra/database"
	"github.com/garcia-paulo/go-gin/Infra/repositories"
	"github.com/garcia-paulo/go-gin/Presentation/controllers"
	"github.com/garcia-paulo/go-gin/Presentation/routes"
	"github.com/google/wire"
)

func InitializeRoutes() *routes.Routes {
	panic(wire.Build(
		database.NewDatabase,
		repositories.NewStudentRepository,
		servicers.NewStudentServicer,
		controllers.NewStudentController,
		routes.NewRoutes,
	))
}
