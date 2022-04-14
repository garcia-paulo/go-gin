package models

import (
	input_student "github.com/garcia-paulo/go-gin/application/dtos/student/input"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string
	CPF  string // document used for citizenship identification in Brazil
}

func NewStudent(student input_student.StudentRequest) *Student {
	return &Student{
		Name: student.Name,
		CPF:  student.CPF,
	}
}

func (s *Student) Validate() error {
	err := validator.Validate(s)
	if err != nil {
		return err
	}
	return nil
}
