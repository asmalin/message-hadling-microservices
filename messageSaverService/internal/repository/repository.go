package repository

import (
	"database/sql"
	"message-saver/internal/model"
)

type Message interface {
	GetTotalMessages() (int, error)
	GetProcessedMessages() (int, error)
	SaveMessage(message model.Message) (msg model.Message, err error)
}

type Repository struct {
	Message
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Message: NewMessageRepo(db),
	}
}
