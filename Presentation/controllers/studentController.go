package controllers

import (
	"net/http"

	"github.com/garcia-paulo/go-gin/Application/servicers"
	"github.com/garcia-paulo/go-gin/Domain/models"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	studentServicer *servicers.StudentServicer
}

func NewStudentController(studentServicer *servicers.StudentServicer) *StudentController {
	return &StudentController{
		studentServicer: studentServicer,
	}
}

func (c *StudentController) FindStudents(context *gin.Context) {
	context.JSON(http.StatusOK, c.studentServicer.FindStudents())
}

func (c *StudentController) FindStudentById(context *gin.Context) {
	id := context.Param("studentId")

	student := c.studentServicer.FindStudentById(id)
	if student.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found.",
		})
		return
	}

	context.JSON(http.StatusOK, student)
}

func (c *StudentController) FindStudentByCpf(context *gin.Context) {
	cpf := context.Param("studentCpf")

	student := c.studentServicer.FindStudentByCpf(cpf)
	if student.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found.",
		})
		return
	}

	context.JSON(http.StatusOK, student)
}

func (c *StudentController) CreateStudent(context *gin.Context) {
	student := models.Student{}

	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := student.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.studentServicer.CreateStudent(&student)
	context.JSON(http.StatusOK, student)
}

func (c *StudentController) UpdateStudent(context *gin.Context) {
	studentId := context.Param("studentId")

	data := models.Student{}
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := data.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	student := c.studentServicer.FindStudentById(studentId)
	c.studentServicer.UpdateStudent(&student, data)
	if student.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found.",
		})
		return
	}

	context.JSON(http.StatusOK, student)
}

func (c *StudentController) DeleteStudent(context *gin.Context) {
	studentId := context.Param("studentId")

	c.studentServicer.DeleteStudent(studentId)
	context.JSON(http.StatusOK, gin.H{
		"message": "Student succesfully deleted.",
	})
}
