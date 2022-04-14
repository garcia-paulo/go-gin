package controllers

import (
	"net/http"

	input_student "github.com/garcia-paulo/go-gin/application/dtos/student/input"
	"github.com/garcia-paulo/go-gin/application/servicers"
	"github.com/garcia-paulo/go-gin/application/utils"
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

	student, err := c.studentServicer.FindStudentById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, student)
}

func (c *StudentController) FindStudentByCpf(context *gin.Context) {
	cpf := context.Param("studentCpf")

	student, err := c.studentServicer.FindStudentByCpf(cpf)
	if err != nil {
		context.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, student)
}

func (c *StudentController) CreateStudent(context *gin.Context) {
	student := input_student.StudentRequest{}

	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	if err := student.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	response, err := c.studentServicer.CreateStudent(student)
	if err != nil {
		context.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}
	context.JSON(http.StatusOK, response)
}

func (c *StudentController) UpdateStudent(context *gin.Context) {
	studentId := context.Param("studentId")

	data := input_student.StudentRequest{}
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
	}
	if err := data.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	response, err := c.studentServicer.UpdateStudent(studentId, data)
	if err != nil {
		context.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, response)
}

func (c *StudentController) DeleteStudent(context *gin.Context) {
	studentId := context.Param("studentId")

	c.studentServicer.DeleteStudent(studentId)
	context.JSON(http.StatusOK, utils.DefaultResponse("student succesfully deleted"))
}
