package controllers

import (
	"net/http"

	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/repositories"
	"github.com/gin-gonic/gin"
)

func FindStudents(c *gin.Context) {
	c.JSON(http.StatusOK, repositories.FindStudents())
}

func FindStudentById(c *gin.Context) {
	id := c.Param("studentId")

	student := repositories.FindStudentById(id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found.",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func FindStudentByCpf(c *gin.Context) {
	cpf := c.Param("studentCpf")

	student := repositories.FindStudentByCpf(cpf)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found.",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := student.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	repositories.CreateStudent(&student)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	studentId := c.Param("studentId")

	data := models.Student{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := data.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	student := repositories.FindStudentById(studentId)
	repositories.UpdateStudent(&student, data)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found.",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	studentId := c.Param("studentId")

	repositories.DeleteStudent(studentId)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student succesfully deleted.",
	})
}
