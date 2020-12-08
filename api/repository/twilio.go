package repository

import (
	"fmt"
	"locator-api/models"
	"locator-api/utils"

	"github.com/kevinburke/twilio-go"
	"gorm.io/gorm"
)

// MessageRepository is ...
type MessageRepository interface {
	SendMessage(*models.Message) (*models.Message, error)
	Migrate() error
}
type messageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{
		DB: db,
	}
}
func (m *messageRepository) SendMessage(message *models.Message) (*models.Message, error) {
	get := utils.GetEnvWithKey
	var userMsg models.Message
	fmt.Println("msg", &message)
	sid := get("ACCOUNT_SID")
	authToken := get("TWILIO_TOKEN")
	client := twilio.NewClient(sid, authToken, nil)
	if client != nil {
		msg, err := client.Messages.SendMessage("+19293251726", "+9779813228397", message.SMS, nil)
		if err != nil && msg != nil {
			fmt.Println("message sent", userMsg)
			m.DB.Create(&message)
			return nil, err
		}

	}

	return nil, nil
}
func (ms *messageRepository) Migrate() error {
	err := ms.DB.AutoMigrate(&models.Message{})
	if err != nil {
		fmt.Println("error", err)
	}
	return nil

}
