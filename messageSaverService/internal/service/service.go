package service

import (
	"message-saver/internal/model"

	"message-saver/internal/repository"
)

type Message interface {
	GetTotalMessages() (int, error)
	GetProcessedMessages() (int, error)
	SaveMessage(message model.MessageInput) (messageId int, err error)
}

type Service struct {
	Message
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repo.Message),
	}
}
