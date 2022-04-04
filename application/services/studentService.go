package services

import (
	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/repositories"
)

func FindStudents() []models.Student {
	return repositories.FindStudents()
}

func CreateStudent(student *models.Student) {
	repositories.CreateStudent(student)
}

func FindStudentById(studentId string) models.Student {
	return repositories.FindStudentById(studentId)
}

func FindStudentByCpf(studentCpf string) models.Student {
	return repositories.FindStudentByCpf(studentCpf)
}

func UpdateStudent(id string, data models.Student) models.Student {
	student := repositories.FindStudentById(id)
	if student.ID == 0 {
		return student
	}

	return repositories.UpdateStudent(student, data)
}

func DeleteStudent(id string) {
	student := repositories.FindStudentById(id)
	repositories.DeleteStudent(student)
}
