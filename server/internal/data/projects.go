package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aravindmathradan/tema/internal/validator"
)

type Project struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Version     int32     `json:"version"`
}

func ValidateProject(v *validator.Validator, project *Project) {
	v.Check(validator.NotBlank(project.Name), "title", validator.EBLANKFIELD, "must be provided")
	v.Check(validator.MaxChars(project.Name, 40), "title", validator.EMAXCHARS, "must not be more than 40 characters long")

	v.Check(validator.MaxBytes(project.Description, 500), "description", validator.EMAXCHARS, "must not be more than 500 bytes long")
}

type ProjectModel struct {
	DB *sql.DB
}

func (m ProjectModel) Get(id int64) (*Project, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, updated_at, name, description, version
		FROM projects
		WHERE id=$1`

	var project Project

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&project.ID,
		&project.CreatedAt,
		&project.UpdatedAt,
		&project.Name,
		&project.Description,
		&project.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &project, nil
}

func (m ProjectModel) Insert(project *Project) error {
	query := `
		INSERT INTO projects (name, description)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at, version`

	args := []any{project.Name, project.Description}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt, &project.Version)
}

func (m ProjectModel) Update(project *Project) error {
	query := `
        UPDATE projects 
        SET name = $1, description = $2, updated_at = $3, version = version + 1
        WHERE id = $4 AND version = $5
		RETURNING version`

	// Create an args slice containing the values for the placeholder parameters.
	args := []any{
		project.Name,
		project.Description,
		time.Now(),
		project.ID,
		project.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&project.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m ProjectModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM projects
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
