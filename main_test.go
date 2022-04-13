package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/garcia-paulo/go-gin/application/servicers"
	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/config"
	"github.com/garcia-paulo/go-gin/infra/database"
	"github.com/garcia-paulo/go-gin/infra/repositories"
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Setup() *gin.Engine {
	duration, _ := time.ParseDuration("1m")
	db = database.NewDatabase(&config.Config{
		DBSource:      "",
		ServerPort:    ":8080",
		TokenKey:      "STRING_NÃO_ESCONDIDA_DE_32_BYTE",
		TokenDuration: duration,
	})
	studentRepository = *repositories.NewStudentRepository(db)
	studentServicer = *servicers.NewStudentServicer(&studentRepository)
	studentController = *controllers.NewStudentController(&studentServicer)
	return gin.Default()
}

var (
	mockStudent       models.Student
	db                *gorm.DB
	studentController controllers.StudentController
	studentServicer   servicers.StudentServicer
	studentRepository repositories.StudentRepository
)

func CreateStudentMock() {
	student := models.Student{
		Name: "Test",
		CPF:  "12345678901",
	}
	db.Create(&student)
	mockStudent = student
}

func DeleteStudentMock() {
	db.Unscoped().Delete(&mockStudent)
}

func TestFindStudents(t *testing.T) {
	r := Setup()
	r.GET("/students", studentController.FindStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, `FindStudents: StatusCodeError; Expected: "%d"; Received: "%d";`, http.StatusOK, response.Code)
}

func TestFindStudentById(t *testing.T) {
	r := Setup()
	CreateStudentMock()
	defer DeleteStudentMock()
	r.GET("/students/:studentId", studentController.FindStudentById)
	req, _ := http.NewRequest("GET", "/students/"+strconv.Itoa(int(mockStudent.ID)), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, `FindStudentsById: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
}

func TestFindStudentByCpf(t *testing.T) {
	r := Setup()
	CreateStudentMock()
	defer DeleteStudentMock()
	r.GET("/students/cpf/:studentCpf", studentController.FindStudentByCpf)
	req, _ := http.NewRequest("GET", "/students/cpf/"+mockStudent.CPF, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, `FindStudentsByCpf: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
}

func TestCreateStudent(t *testing.T) {
	r := Setup()
	defer DeleteStudentMock()
	r.POST("/students", studentController.CreateStudent)

	data := models.Student{
		Name: "Test",
		CPF:  "12345678901",
	}

	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	createdStudent := models.Student{}
	json.Unmarshal(response.Body.Bytes(), &createdStudent)
	mockStudent = createdStudent

	assert.Equal(t, http.StatusOK, response.Code, `FindStudents: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
}

func TestUpdateStudent(t *testing.T) {
	r := Setup()
	CreateStudentMock()
	defer DeleteStudentMock()
	r.PATCH("/students/:studentId", studentController.UpdateStudent)
	data := models.Student{
		Name: "Update Test",
		CPF:  "12345678901",
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("PATCH", "/students/"+strconv.Itoa(int(mockStudent.ID)), bytes.NewBuffer(jsonData))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	updatedStudent := models.Student{}
	json.Unmarshal(response.Body.Bytes(), &updatedStudent)

	assert.Equal(t, http.StatusOK, response.Code, `TestUpdateStudent: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
	assert.Equal(t, data.Name, updatedStudent.Name, `TestUpdateStudent: FieldValueError; Expected: "%s"; Received: "%s"`, data.Name, updatedStudent.Name)
}

func TestDeleteStudent(t *testing.T) {
	r := Setup()
	CreateStudentMock()
	defer DeleteStudentMock()
	r.DELETE("/students/:studentId", studentController.DeleteStudent)
	req, _ := http.NewRequest("DELETE", "/students/"+strconv.Itoa(int(mockStudent.ID)), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, `DeleteStudent: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
}
