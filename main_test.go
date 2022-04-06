package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/database"
	"github.com/garcia-paulo/go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func RouteTestSetup() *gin.Engine {
	return gin.Default()
}

var mockStudent models.Student

func CreateStudentMock() {
	student := models.Student{
		Name: "Test",
		CPF:  "12345678901",
		RG:   "123456789",
	}
	database.DB.Create(&student)
	mockStudent = student
}

func DeleteStudentMock() {
	database.DB.Unscoped().Delete(&mockStudent)
}

func TestFindStudents(t *testing.T) {
	database.DBConnect()
	r := RouteTestSetup()
	r.GET("/students", controllers.FindStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, `FindStudents: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
}

func TestFindStudentById(t *testing.T) {
	database.DBConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouteTestSetup()
	r.GET("/students/:studentId", controllers.FindStudentById)
	req, _ := http.NewRequest("GET", "/students/"+strconv.Itoa(int(mockStudent.ID)), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, `FindStudentsById: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
}

func TestFindStudentByCpf(t *testing.T) {
	database.DBConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouteTestSetup()
	r.GET("/students/cpf/:studentCpf", controllers.FindStudentByCpf)
	req, _ := http.NewRequest("GET", "/students/cpf/"+mockStudent.CPF, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, `FindStudentsByCpf: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
}

func TestCreateStudent(t *testing.T) {
	database.DBConnect()
	defer DeleteStudentMock()
	r := RouteTestSetup()
	r.POST("/students", controllers.CreateStudent)

	data := models.Student{
		Name: "Test",
		CPF:  "12345678901",
		RG:   "123456789",
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
	database.DBConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouteTestSetup()
	r.PATCH("/students/:studentId", controllers.UpdateStudent)
	data := models.Student{
		Name: "Update Test",
		CPF:  "12345678901",
		RG:   "123456789",
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
	database.DBConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RouteTestSetup()
	r.DELETE("/students/:studentId", controllers.DeleteStudent)
	req, _ := http.NewRequest("DELETE", "/students/"+strconv.Itoa(int(mockStudent.ID)), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, `DeleteStudent: StatusCodeError; Expected: "%d"; Received: "%d"`, http.StatusOK, response.Code)
}
