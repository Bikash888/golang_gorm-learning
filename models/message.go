package models

import (
	"github.com/jinzhu/gorm"
)

//Message model
type Message struct {
	gorm.Model
	message string
}
