package routes

import (
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
)

type StudentRoutes struct {
	studentController *controllers.StudentController
}

func NewRoutes(studentController *controllers.StudentController) *StudentRoutes {
	return &StudentRoutes{
		studentController: studentController,
	}
}

func (r *StudentRoutes) HandleRequests() {
	s := gin.Default()
	s.GET("/students", r.studentController.FindStudents)
	s.GET("/students/:studentId", r.studentController.FindStudentById)
	s.GET("/students/search/:studentCpf", r.studentController.FindStudentByCpf)
	s.POST("/students", r.studentController.CreateStudent)
	s.PATCH("/students/:studentId", r.studentController.UpdateStudent)
	s.DELETE("/students/:studentId", r.studentController.DeleteStudent)
	s.Run()
}
