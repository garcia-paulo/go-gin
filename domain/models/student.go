package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"min=3,max=40,regexp=^[a-zA-Z\\s]+$"`
	CPF  string `json:"cpf" validate:"len=11,regexp=^[0-9]+$" gorm:"uniqueIndex"` // document used for citizenship identification in Brazil
}

func (student *Student) Validate() error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
