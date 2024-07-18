package repository

import (
	"database/sql"
	"fmt"
	"message-saver/internal/model"
)

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) GetTotalMessages() (int, error) {

	var messageCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&messageCount)
	if err != nil {
		return 0, err
	}

	return messageCount, nil
}

func (r *MessageRepo) GetProcessedMessages() (int, error) {
	var messageCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM messages WHERE processed = true").Scan(&messageCount)
	if err != nil {
		return 0, err
	}

	return messageCount, nil
}

func (r *MessageRepo) SaveMessage(message model.Message) (msg model.Message, err error) {
	query := `
		INSERT INTO messages (text)
		VALUES ($1)
		RETURNING id, text, processed, created_at`

	err = r.db.QueryRow(query, message.Text).Scan(&msg.Id, &msg.Text, &msg.Processed, &msg.CreatedAt)
	if err != nil {
		return model.Message{}, fmt.Errorf("failed to create message: %w", err)
	}

	return msg, nil
}
