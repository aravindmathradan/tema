package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/aravindmathradan/tema/internal/validator"
)

type Transcript struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Archived    bool      `json:"archived"`
	ProjectID   int64     `json:"project_id"`
	Version     int32     `json:"version"`
}

func ValidateTranscript(v *validator.Validator, transcript *Transcript) {
	v.Check(validator.NotBlank(transcript.Title), "title", validator.EBLANKFIELD, "must be provided")
	v.Check(validator.MaxChars(transcript.Title, 40), "title", validator.EMAXCHARS, "must not be more than 40 characters long")

	v.Check(validator.MaxBytes(transcript.Description, 500), "description", validator.EMAXCHARS, "must not be more than 500 bytes long")

	v.Check(validator.NotBlank(transcript.Content), "content", validator.EBLANKFIELD, "must be provided")
}

type TranscriptModel struct {
	DB *sql.DB
}

func (m TranscriptModel) Get(id int64) (*Transcript, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, updated_at, title, description, content, archived, project_id, version
		FROM transcripts
		WHERE id = $1`

	var transcript Transcript

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&transcript.ID,
		&transcript.CreatedAt,
		&transcript.UpdatedAt,
		&transcript.Title,
		&transcript.Description,
		&transcript.Content,
		&transcript.Archived,
		&transcript.Archived,
		&transcript.ProjectID,
		&transcript.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &transcript, nil
}

func (m TranscriptModel) GetAllForProject(projectID int64, title string, archived bool, filters Filters) ([]*Transcript, Metadata, error) {
	query := fmt.Sprintf(`SELECT COUNT(*) OVER(), id, created_at, updated_at, title, description, content, archived, project_id, version
	FROM transcripts
	WHERE project_id = $1 AND archived = $2
	AND (to_tsvector('english', title) @@ plainto_tsquery('english', $3) OR  $3 = '')
	ORDER BY %s %s, id ASC
	LIMIT $4 OFFSET $5`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{projectID, archived, title, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	transcripts := []*Transcript{}

	for rows.Next() {
		var transcript Transcript

		err := rows.Scan(
			&totalRecords,
			&transcript.ID,
			&transcript.CreatedAt,
			&transcript.UpdatedAt,
			&transcript.Title,
			&transcript.Description,
			&transcript.Content,
			&transcript.Archived,
			&transcript.ProjectID,
			&transcript.Version,
		)

		if err != nil {
			return nil, Metadata{}, err
		}

		transcripts = append(transcripts, &transcript)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return transcripts, metadata, nil
}

func (m TranscriptModel) Insert(transcript *Transcript) error {
	query := `
	INSERT INTO transcripts (title, description, content, project_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at, archived, version`

	args := []any{transcript.Title, transcript.Description, transcript.Content, transcript.ProjectID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&transcript.ID,
		&transcript.CreatedAt,
		&transcript.UpdatedAt,
		&transcript.Archived,
		&transcript.Version,
	)
}

func (m TranscriptModel) Update(transcript *Transcript) error {
	query := `
        UPDATE transcripts 
        SET title = $1, description = $2, content = $3, archived = $4, updated_at = $5, version = version + 1
        WHERE id = $6 AND version = $7
		RETURNING version`

	// Create an args slice containing the values for the placeholder parameters.
	args := []any{
		transcript.Title,
		transcript.Description,
		transcript.Content,
		transcript.Archived,
		time.Now(),
		transcript.ID,
		transcript.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&transcript.Version)
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

func (m TranscriptModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM transcripts
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
