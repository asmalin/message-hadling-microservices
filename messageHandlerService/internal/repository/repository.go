package repository

import (
	"database/sql"
)

type Message interface {
	FlagMessage(messageID int) error
}

type Repository struct {
	Message
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Message: NewMessageRepo(db),
	}
}
