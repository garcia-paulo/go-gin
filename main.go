package main

import (
	"github.com/garcia-paulo/go-gin/infra/database"
	"github.com/garcia-paulo/go-gin/presentation/routes"
)

func main() {
	database.DBConnect()
	routes.HandleRequests()
}
