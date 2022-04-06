package repositories

import (
	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/database"
)

func FindStudents() []models.Student {
	students := []models.Student{}
	database.DB.Find(&students)
	return students
}

func FindStudentById(id string) models.Student {
	student := models.Student{}
	database.DB.First(&student, id)
	return student
}

func FindStudentByCpf(cpf string) models.Student {
	student := models.Student{}
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	return student
}

func CreateStudent(student *models.Student) {
	database.DB.Create(&student)
}

func UpdateStudent(student *models.Student, data models.Student) {
	database.DB.Model(&student).UpdateColumns(data)
}

func DeleteStudent(studentId string) {
	database.DB.Delete(&models.Student{}, studentId)
}
