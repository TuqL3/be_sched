package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func ConversationRoute(route *gin.Engine, controller *controllers.ConversationController) {
	conversationGroup := route.Group("/api/v1/conversation")
	{
		conversationGroup.POST("/findorcreate", controller.FindOrCreateConversation)
		conversationGroup.GET("/", controller.GetAllConversation)
		conversationGroup.GET("/:conversationId", controller.GetConversaionById)
		conversationGroup.PUT("/:conversationId", controller.UpdateLastMessage)
	}
}
