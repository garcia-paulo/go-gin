package routes

import "github.com/gin-gonic/gin"

type Server struct {
	studentRoutes *StudentRoutes
}

func NewServer(studentRoutes *StudentRoutes) *Server {
	return &Server{
		studentRoutes: studentRoutes,
	}
}

func (s *Server) Serve() {
	g := gin.Default()
	s.studentRoutes.HandleRequests(g)
	g.Run()
}
