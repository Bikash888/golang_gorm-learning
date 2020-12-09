package repository

import (
	"locator-api/models"
)

//PostRepository
type PostRepository interface {
	FindPostById(id int) (*models.Posts, error)
	Migrate() error
}

type postRepository struct {
}
