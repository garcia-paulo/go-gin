//go:build wireinject
// +build wireinject

package main

import (
	"github.com/garcia-paulo/go-gin/application/servicers"
	"github.com/garcia-paulo/go-gin/infra/database"
	"github.com/garcia-paulo/go-gin/infra/repositories"
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/garcia-paulo/go-gin/presentation/routes"
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
