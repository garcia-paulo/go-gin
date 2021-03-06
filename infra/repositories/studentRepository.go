package repositories

import (
	"fmt"

	"github.com/garcia-paulo/go-gin/domain/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	database *gorm.DB
}

func NewStudentRepository(database *gorm.DB) *StudentRepository {
	return &StudentRepository{
		database: database,
	}
}

func (r *StudentRepository) FindStudents() []models.Student {
	students := []models.Student{}
	r.database.Find(&students)
	return students
}

func (r *StudentRepository) FindStudentById(studentId string) (*models.Student, error) {
	student := &models.Student{}
	r.database.First(&student, studentId)
	if student.ID == 0 {
		return nil, fmt.Errorf("student not found")
	}
	return student, nil
}

func (r *StudentRepository) FindStudentByCpf(cpf string) (*models.Student, error) {
	student := &models.Student{}
	r.database.Where(models.Student{CPF: cpf}).First(student)
	if student.ID == 0 {
		return nil, fmt.Errorf("student not found")
	}
	return student, nil
}

func (r *StudentRepository) CreateStudent(student *models.Student) {
	r.database.Create(&student)
}

func (r *StudentRepository) UpdateStudent(student *models.Student, data models.Student) {
	r.database.Model(&student).UpdateColumns(data)
}

func (r *StudentRepository) DeleteStudent(studentId string) {
	r.database.Delete(&models.Student{}, studentId)
}
