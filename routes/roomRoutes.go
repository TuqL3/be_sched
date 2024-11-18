package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func RoomRoute(route *gin.Engine, controller *controllers.RoomController) {
	roomRouteMiddleware := route.Group("/api/v1/room")
	{
		roomRouteMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"createRoom"}), controller.CreateRoom)
		roomRouteMiddleware.DELETE("/delete/:roomId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"deleteRoom"}), controller.DeleteRoom)
		roomRouteMiddleware.PUT("/update/:roomId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"modifyRoom"}), controller.UpdateRoom)
		roomRouteMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"viewRoom"}), controller.GetAllRoom)
		roomRouteMiddleware.GET("/getcountroom", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"viewRoom"}), controller.GetCountRoom)
		roomRouteMiddleware.GET("/:roomId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"viewRoom"}), controller.GetRoomById)
	}
}
