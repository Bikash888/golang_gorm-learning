package models

import "github.com/jinzhu/gorm"

//Comments ..
type Comments struct {
	gorm.Model
	Comment   string `json:"comment"`
	UserRefer uint
}
