package routes

import (
	"fmt"
	"locator-api/api/controllers"
	"locator-api/api/repository"
	"locator-api/api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//PostRoutes ..
func PostRoutes(routes *gin.RouterGroup, db *gorm.DB) {
	postRepo := repository.NewPostRepository(db)
	if err := postRepo.Migrate(); err != nil {
		fmt.Println("Migration err--->", err)
	}
	postService := service.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)
	routes.GET("/:id", postController.FindPostByID)
}
