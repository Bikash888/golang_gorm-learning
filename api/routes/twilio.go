package routes

import (
	"fmt"
	"locator-api/api/controllers"
	"locator-api/api/repository"
	"locator-api/api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TwilioRoutes ..
func TwilioRoutes(route *gin.RouterGroup, db *gorm.DB) {
	smsRepository := repository.NewMessageRepository(db)
	if err := smsRepository.Migrate(); err != nil {
		fmt.Println(" migrate err", err)
	}
	smsService := service.NewMessageService(smsRepository)
	smsController := controllers.NewMessageController(smsService)
	route.POST("/sms", smsController.SaveAndSendMessage)

}
