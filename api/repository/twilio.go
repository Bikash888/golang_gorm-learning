package repository

import (
	"locator-api/models"

	"gorm.io/gorm"
)

// MessageRepository is ...
type MessageRepository interface {
	SendMessage(*models.Message) (*models.Message, error)
}
type messageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{
		DB: db,
	}
}
func (m *messageRepository) SendMessage(*models.Message) (*models.Message, error) {

	return nil, nil
}
