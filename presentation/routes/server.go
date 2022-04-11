package routes

import "github.com/gin-gonic/gin"

type Server struct {
	studentRoutes *StudentRoutes
	userRoutes    *UserRoutes
}

func NewServer(studentRoutes *StudentRoutes, userRoutes *UserRoutes) *Server {
	return &Server{
		studentRoutes: studentRoutes,
		userRoutes:    userRoutes,
	}
}

func (s *Server) Serve() {
	g := gin.Default()
	s.studentRoutes.HandleRequests(g)
	s.userRoutes.HandleRequests(g)
	g.Run()
}
