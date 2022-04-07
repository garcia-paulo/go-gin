package routes

import (
	"github.com/garcia-paulo/go-gin/Presentation/controllers"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	studentController *controllers.StudentController
}

func NewRoutes(studentController *controllers.StudentController) *Routes {
	return &Routes{
		studentController: studentController,
	}
}

func (r *Routes) HandleRequests() {
	s := gin.Default()
	s.GET("/students", r.studentController.FindStudents)
	s.GET("/students/:studentId", r.studentController.FindStudentById)
	s.GET("/students/search/:studentCpf", r.studentController.FindStudentByCpf)
	s.POST("/students", r.studentController.CreateStudent)
	s.PATCH("/students/:studentId", r.studentController.UpdateStudent)
	s.DELETE("/students/:studentId", r.studentController.DeleteStudent)
	s.Run()
}
