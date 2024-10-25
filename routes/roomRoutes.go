package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func RoomRoute(route *gin.Engine, controller *controllers.RoomController) {
	roomRouteMiddleware := route.Group("/api/v1/room")
	{
		roomRouteMiddleware.Use(middleware.AuthMiddleware())
		roomRouteMiddleware.Use(middleware.AdminOnly())

		roomRouteMiddleware.POST("/create", controller.CreateRoom)
		roomRouteMiddleware.DELETE("/delete/:roomId", controller.DeleteRoom)
		roomRouteMiddleware.PUT("/update/:roomId", controller.UpdateRoom)
		roomRouteMiddleware.GET("", controller.GetAllRoom)
	}
}
