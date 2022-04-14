package servicers

import (
	"fmt"
	"time"

	input_student "github.com/garcia-paulo/go-gin/application/dtos/student/input"
	output_student "github.com/garcia-paulo/go-gin/application/dtos/student/output"
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

func (s *StudentServicer) FindStudents() []output_student.StudentResponse {
	students := s.studentRepository.FindStudents()
	response := []output_student.StudentResponse{}
	for _, s := range students {
		response = append(response, *output_student.NewStudentResponse(&s))
	}
	return response
}

func (s *StudentServicer) FindStudentById(studentId string) (*output_student.StudentResponse, error) {
	student, err := s.studentRepository.FindStudentById(studentId)
	if err != nil {
		return nil, err
	}
	return output_student.NewStudentResponse(student), nil
}

func (s *StudentServicer) FindStudentByCpf(cpf string) (*output_student.StudentResponse, error) {
	student, err := s.studentRepository.FindStudentByCpf(cpf)
	if err != nil {
		return nil, err
	}
	return output_student.NewStudentResponse(student), nil
}

func (s *StudentServicer) CreateStudent(data input_student.StudentRequest) (*output_student.StudentResponse, error) {
	student := models.NewStudent(data)
	s.studentRepository.CreateStudent(student)
	if student.ID == 0 {
		return nil, fmt.Errorf("error when saving to the database")
	}
	return output_student.NewStudentResponse(student), nil
}

func (s *StudentServicer) UpdateStudent(studentId string, data input_student.StudentRequest) (*output_student.StudentResponse, error) {
	student, err := s.studentRepository.FindStudentById(studentId)
	if err != nil {
		return nil, fmt.Errorf("student not found")
	}
	updatedStudent := models.NewStudent(data)
	updatedStudent.UpdatedAt = time.Now()
	s.studentRepository.UpdateStudent(student, *updatedStudent)
	return output_student.NewStudentResponse(student), nil
}

func (s *StudentServicer) DeleteStudent(studentId string) {
	s.studentRepository.DeleteStudent(studentId)
}
