package repository

import (
	"database/sql"
)

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) FlagMessage(messageID int) error {

	query := "UPDATE messages SET processed = true WHERE id = $1"

	err := r.db.QueryRow(query, messageID)
	if err != nil {
		return err.Err()
	}

	return nil
}
