package input_student

import (
	"gopkg.in/validator.v2"
)

type StudentRequest struct {
	Name string `json:"name" validate:"min=3,max=40,regexp=^[a-zA-Z\\s]+$"`
	CPF  string `json:"cpf" validate:"len=11,regexp=^[0-9]+$" gorm:"uniqueIndex"`
}

func (student *StudentRequest) Validate() error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
