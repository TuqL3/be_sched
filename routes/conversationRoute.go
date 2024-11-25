package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func ConversationRoute(route *gin.Engine, controller *controllers.ConversationController) {
	conversationGroup := route.Group("/api/v1/conversation")
	{
		conversationGroup.POST("/findorcreate", controller.FindOrCreateConversation)
		conversationGroup.GET("/", controller.GetAllConversation)
		conversationGroup.GET("/:conversationId", controller.GetConversaionById)
		conversationGroup.POST("/:conversationId", controller.UpdateLastMessage)
	}
}
