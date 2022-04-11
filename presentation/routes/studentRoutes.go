package routes

import (
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
)

type StudentRoutes struct {
	studentController *controllers.StudentController
}

func NewStudentRoutes(studentController *controllers.StudentController) *StudentRoutes {
	return &StudentRoutes{
		studentController: studentController,
	}
}

func (r *StudentRoutes) HandleRequests(g *gin.Engine) {
	g.GET("/students", r.studentController.FindStudents)
	g.GET("/students/:studentId", r.studentController.FindStudentById)
	g.GET("/students/search/:studentCpf", r.studentController.FindStudentByCpf)
	g.POST("/students", r.studentController.CreateStudent)
	g.PATCH("/students/:studentId", r.studentController.UpdateStudent)
	g.DELETE("/students/:studentId", r.studentController.DeleteStudent)
}
