package output_student

import (
	"time"

	"github.com/garcia-paulo/go-gin/domain/models"
	"gorm.io/gorm"
)

type StudentResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	CPF       string         `json:"cpf"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func NewStudentResponse(student *models.Student) *StudentResponse {
	return &StudentResponse{
		ID:        student.ID,
		Name:      student.Name,
		CPF:       student.CPF,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
		DeletedAt: student.DeletedAt,
	}
}
