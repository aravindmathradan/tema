package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Projects ProjectModel
	Tokens   TokenModel
	Users    UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Projects: ProjectModel{DB: db},
		Tokens:   TokenModel{DB: db},
		Users:    UserModel{DB: db},
	}
}
