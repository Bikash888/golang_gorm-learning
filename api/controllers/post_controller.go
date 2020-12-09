package controllers

import (
	"locator-api/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//PostController ..
type PostController interface {
	FindPostByID(c *gin.Context)
}

type postController struct {
	service service.PostService
}

func NewPostController(service service.PostService) PostController {
	return &postController{
		service: service,
	}
}

func (pc *postController) FindPostByID(c *gin.Context) {
	listPosts, err := pc.service.FindPostDataByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, listPosts)
}
