package service

import (
	"message-saver/internal/kafka"
	"message-saver/internal/model"
	"message-saver/internal/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) GetTotalMessages() (int, error) {
	total, err := s.repo.GetTotalMessages()
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s *MessageService) GetProcessedMessages() (int, error) {

	processedMsg, err := s.repo.GetProcessedMessages()
	if err != nil {
		return 0, err
	}

	return processedMsg, nil
}

func (s *MessageService) SaveMessage(message model.MessageInput) (msg model.Message, err error) {

	msg.Text = message.Text
	msg.Processed = false

	msg, err = s.repo.SaveMessage(msg)

	if err != nil {
		return model.Message{}, err
	}

	kafka.SendMessage(msg)

	return msg, nil
}
