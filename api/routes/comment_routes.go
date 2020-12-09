package routes

import (
	"fmt"
	"locator-api/api/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CommentRoutes ..
func CommentRoutes(routes *gin.RouterGroup, db *gorm.DB) {
	commentRepo := repository.NewCommentRepository(db)
	if err := commentRepo.Migrate(); err != nil {
		fmt.Println("Migration Error===>", err)
	}
}
