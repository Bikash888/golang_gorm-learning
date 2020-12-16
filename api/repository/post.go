package repository

import (
	"fmt"
	"locator-api/models"

	"gorm.io/gorm"
)

//PostRepository ..
type PostRepository interface {
	FindPostByID(id int) (*models.Posts, error)
	Migrate() error
}

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		DB: db,
	}
}
func (p *postRepository) FindPostByID(id int) (*models.Posts, error) {
	var post models.Posts
	result := p.DB.Where("id=?", id).Preload("Comment").First(&post)
	if result.Error != nil {
		err := result.Error
		fmt.Println("Error>>", err)
		return nil, err
	}
	return &post, nil

}
func (p *postRepository) Migrate() error {
	err := p.DB.AutoMigrate(&models.Posts{})
	if err != nil {
		fmt.Println("error", err)
	}
	return nil

}
