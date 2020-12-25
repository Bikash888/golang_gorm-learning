package repository

import (
	"fmt"
	"locator-api/models"

	"gorm.io/gorm"
)

//PostRepository ..
type PostRepository interface {
	FindPostByID(id int) (*models.Posts, error)
	FindAllPost(post *models.Posts, pagination *models.Pagination) (*[]models.Posts, int64, error)
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

func (p *postRepository) FindAllPost(post *models.Posts, pagination *models.Pagination) (*[]models.Posts, int64, error) {
	var posts []models.Posts
	var totalRow int64 = 0
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Posts{}).Where(post).Find(&posts)
	if result.Error != nil {
		msg := result.Error
		return nil, totalRow, msg
	}
	return &posts, totalRow, nil
}
