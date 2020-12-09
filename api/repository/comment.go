package repository

import (
	"fmt"
	"locator-api/models"

	"gorm.io/gorm"
)

// CommentRepository ..
type CommentRepository interface {
	Migrate() error
}

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		DB: db,
	}

}

func (p *commentRepository) Migrate() error {
	err := p.DB.AutoMigrate(&models.Comments{})
	if err != nil {
		fmt.Println("error", err)
	}
	return nil

}
