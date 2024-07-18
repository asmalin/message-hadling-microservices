package service

import (
	"message-handler/internal/model"
	"message-handler/internal/repository"
)

type Message interface {
	ProcessMessage(message model.Message) error
}

type Service struct {
	Message
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repo.Message),
	}
}
