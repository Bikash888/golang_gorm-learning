package models

import (
	"github.com/jinzhu/gorm"
)

//Message model
type Message struct {
	gorm.Model
	Message string `json:"message"`
}
