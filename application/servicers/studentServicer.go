package servicers

import (
	"time"

	"github.com/garcia-paulo/go-gin/domain/models"
	"github.com/garcia-paulo/go-gin/infra/repositories"
)

type StudentServicer struct {
	studentRepository *repositories.StudentRepository
}

func NewStudentServicer(studentRepository *repositories.StudentRepository) *StudentServicer {
	return &StudentServicer{
		studentRepository: studentRepository,
	}
}

func (s *StudentServicer) FindStudents() []models.Student {
	return s.studentRepository.FindStudents()
}

func (s *StudentServicer) FindStudentById(studentId string) models.Student {
	return s.studentRepository.FindStudentById(studentId)
}

func (s *StudentServicer) FindStudentByCpf(cpf string) models.Student {
	return s.studentRepository.FindStudentByCpf(cpf)
}

func (s *StudentServicer) CreateStudent(student *models.Student) {
	s.studentRepository.CreateStudent(student)
}

func (s *StudentServicer) UpdateStudent(student *models.Student, data models.Student) {
	student.UpdatedAt = time.Now()
	s.studentRepository.UpdateStudent(student, data)
}

func (s *StudentServicer) DeleteStudent(studentId string) {
	s.studentRepository.DeleteStudent(studentId)
}
