package routes

import (
	"github.com/garcia-paulo/go-gin/infra/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	studentRoutes *StudentRoutes
	userRoutes    *UserRoutes
	config        *config.Config
}

func NewServer(studentRoutes *StudentRoutes, userRoutes *UserRoutes, config *config.Config) *Server {
	return &Server{
		studentRoutes: studentRoutes,
		userRoutes:    userRoutes,
		config:        config,
	}
}

func (s *Server) Serve() {
	g := gin.Default()
	s.studentRoutes.HandleRequests(g)
	s.userRoutes.HandleRequests(g)
	g.Run(s.config.ServerPort)
}
