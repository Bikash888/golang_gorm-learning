package service

import (
	"fmt"
	"locator-api/api/repository"
	"locator-api/models"
	"locator-api/utils"
)

//PostService ..
type PostService interface {
	FindPostDataByID(id string) (*models.Posts, error)
	FindPostData(posts *models.Posts, pagination *models.Pagination) (*[]models.Posts, int64, error)
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{
		repository: repo,
	}

}

func (ps *postService) FindPostDataByID(id string) (*models.Posts, error) {
	convertedID, err := utils.ConvertStringToInt(id)
	if err != nil {
		fmt.Println("Erorr In converting Id")
	}
	return ps.repository.FindPostByID(convertedID)
}
func (ps *postService) FindPostData(posts *models.Posts, pagination *models.Pagination) (*[]models.Posts, int64, error) {
	return ps.repository.FindAllPost(posts, pagination)
}
