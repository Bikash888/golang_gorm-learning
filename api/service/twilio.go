package service

import (
	"locator-api/api/repository"
	"locator-api/models"
)

// MessageService is ..
type MessageService interface {
	SendMessage(message *models.Message) (*models.Message, error)
}
type messageService struct {
	repository repository.MessageRepository
}

func NewMessageService(msgRepo repository.MessageRepository) MessageService {
	return &messageService{
		repository: msgRepo,
	}
}

func (ms *messageService) SendMessage(message *models.Message) (*models.Message, error) {
	return ms.repository.SendMessage(message)
}
