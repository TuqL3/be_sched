package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func MessageRoute(route *gin.Engine, controller *controllers.MessageController) {
	messageGroup := route.Group("/api/v1/message")
	{
		messageGroup.POST("/", controller.SendMessage)
		messageGroup.GET("/:conversationId", controller.GetMessageByConversationId)
	}
}
