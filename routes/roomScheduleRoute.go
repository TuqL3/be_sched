package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func RoomScheduleRoute(route *gin.Engine, controller *controllers.RoomScheduleController) {
	roomScheduleMiddleware := route.Group("/api/v1/schedule")
	{
		roomScheduleMiddleware.POST("/create", controller.CreateRoomSchedule)
		roomScheduleMiddleware.PUT("/update/:roomScheduleId", controller.UpdateRoomSchedule)
		roomScheduleMiddleware.DELETE("/delete/:roomScheduleId", controller.DeleteRoomSchedule)
		roomScheduleMiddleware.GET("", controller.GetAllRoomSchedule)
	}
}
