package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	CPF  string `json:"cpf"` // document used for citizenship identification in Brazil
	RG   string `json:"rg"`  // document user for general identification inside the country of Brazil
}
