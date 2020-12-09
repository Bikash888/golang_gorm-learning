package models

import "github.com/jinzhu/gorm"

// Posts is ..
type Posts struct {
	gorm.Model
	Status  string     `json:"status"`
	Comment []Comments `gorm:"foreignKey:PostID"`
}
