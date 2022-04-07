package controllers

import (
	"net/http"

	"github.com/garcia-paulo/go-gin/Domain/models"
	"github.com/garcia-paulo/go-gin/Infra/repositories"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	studentRepository *repositories.StudentRepository
}

func NewStudentController(studentRepository *repositories.StudentRepository) *StudentController {
	return &StudentController{
		studentRepository: studentRepository,
	}
}

func (c *StudentController) FindStudents(context *gin.Context) {
	context.JSON(http.StatusOK, c.studentRepository.FindStudents())
}

func (c *StudentController) FindStudentById(context *gin.Context) {
	id := context.Param("studentId")

	student := c.studentRepository.FindStudentById(id)
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

	student := c.studentRepository.FindStudentByCpf(cpf)
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

	c.studentRepository.CreateStudent(&student)
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

	student := c.studentRepository.FindStudentById(studentId)
	c.studentRepository.UpdateStudent(&student, data)
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

	c.studentRepository.DeleteStudent(studentId)
	context.JSON(http.StatusOK, gin.H{
		"message": "Student succesfully deleted.",
	})
}
