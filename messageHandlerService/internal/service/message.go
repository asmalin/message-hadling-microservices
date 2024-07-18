package service

import (
	"message-handler/internal/model"
	"message-handler/internal/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) ProcessMessage(message model.Message) error {

	err := s.repo.FlagMessage(message.Id)
	if err != nil {
		return err
	}

	return nil

}
