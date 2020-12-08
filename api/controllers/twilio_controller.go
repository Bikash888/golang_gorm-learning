package controllers

import (
	"locator-api/api/service"
	"locator-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//MessageController ..
type MessageController interface {
	SaveAndSendMessage(c *gin.Context)
}

type smsController struct {
	messageService service.MessageService
}

func NewMessageController(service service.MessageService) MessageController {
	return &smsController{
		messageService: service,
	}
}

func (sc *smsController) SaveAndSendMessage(c *gin.Context) {
	var messageModel models.Message
	err := c.ShouldBind(&messageModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err,
		})
	}
	sendSms, err := sc.messageService.SendMessage(&messageModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": sendSms,
	})
	return
}
