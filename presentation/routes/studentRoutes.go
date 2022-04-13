package routes

import (
	middleware "github.com/garcia-paulo/go-gin/application/middlewares/auth"
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
)

type StudentRoutes struct {
	studentController *controllers.StudentController
	authMiddleware    *middleware.AuthMiddleware
}

func NewStudentRoutes(studentController *controllers.StudentController, authMiddleware *middleware.AuthMiddleware) *StudentRoutes {
	return &StudentRoutes{
		studentController: studentController,
		authMiddleware:    authMiddleware,
	}
}

func (r *StudentRoutes) HandleRequests(g *gin.Engine) {
	routes := g.Group("/students").Use(r.authMiddleware.Authenticate())

	routes.GET("/", r.studentController.FindStudents)
	routes.GET("/:studentId", r.studentController.FindStudentById)
	routes.GET("/search/:studentCpf", r.studentController.FindStudentByCpf)
	routes.POST("/", r.studentController.CreateStudent)
	routes.PATCH("/:studentId", r.studentController.UpdateStudent)
	routes.DELETE("/:studentId", r.studentController.DeleteStudent)
}
