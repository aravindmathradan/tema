package data

import (
	"time"

	"github.com/aravindmathradan/tema/internal/validator"
)

type Project struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      uint8     `json:"status"`
	UserID      int64     `json:"user_id"`
}

func ValidateProject(v *validator.Validator, project *Project) {
	v.Check(validator.NotBlank(project.Name), "title", "must be provided")
	v.Check(validator.MaxChars(project.Name, 40), "title", "must not be more than 40 characters long")

	v.Check(validator.MaxBytes(project.Description, 500), "description", "must not be more than 500 bytes long")
}
