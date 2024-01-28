package data

import "time"

type Project struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      uint8     `json:"status"`
}
