package infrastructure

import (
	"locator-api/api/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	routes.TwilioRoutes(httpRouter.Group("send"), db)
	routes.PostRoutes(httpRouter.Group("post"), db)
	routes.CommentRoutes(httpRouter.Group("comment"), db)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		httpRouter.Run()
	} else {
		httpRouter.Run(os.Getenv("SERVER_PORT"))
	}

}
