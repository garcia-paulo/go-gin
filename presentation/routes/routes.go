package routes

import (
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.FindStudents)
	r.GET("/students/:studentId", controllers.FindStudentById)
	r.GET("/students/search/:studentCpf", controllers.FindStudentByCpf)
	r.POST("/students", controllers.CreateStudent)
	r.PATCH("/students/:studentId", controllers.UpdateStudent)
	r.DELETE("/students/:studentId", controllers.DeleteStudent)
	r.Run()
}
