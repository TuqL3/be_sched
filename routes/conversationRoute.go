package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func ConversationRoute(route *gin.Engine, controller *controllers.ConversationController) {
	conversationGroup := route.Group("/api/v1/conversation")
	{
		conversationGroup.GET("/findorcreate", controller.FindOrCreateConversation)
		conversationGroup.POST("/:conversationId", controller.UpdateLastMessage)
	}
}
