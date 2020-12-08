package models

import (
	"github.com/jinzhu/gorm"
)

//Twilio model
type Message struct {
	gorm.Model
	message string
}
