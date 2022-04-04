package controllers

import (
	"net/http"

	"github.com/garcia-paulo/go-gin/application/services"
	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/gin-gonic/gin"
)

func FindStudents(c *gin.Context) {
	c.JSON(http.StatusOK, services.FindStudents())
}

func FindStudentById(c *gin.Context) {
	id := c.Param("studentId")
	student := services.FindStudentById(id)

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
	student := services.FindStudentByCpf(cpf)

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
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	services.CreateStudent(&student)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	studentId := c.Param("studentId")
	data := models.Student{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	student := services.UpdateStudent(studentId, data)
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
	services.DeleteStudent(studentId)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student succesfully deleted.",
	})
}
