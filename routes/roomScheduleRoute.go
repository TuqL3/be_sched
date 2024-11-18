package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func RoomScheduleRoute(route *gin.Engine, controller *controllers.RoomScheduleController) {
	roomScheduleMiddleware := route.Group("/api/v1/schedule")
	{
		roomScheduleMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin", "giangvien"}, []string{"create"}), controller.CreateRoomSchedule)
		roomScheduleMiddleware.PUT("/update/:roomScheduleId", controller.UpdateRoomSchedule)
		roomScheduleMiddleware.DELETE("/delete/:roomScheduleId", controller.DeleteRoomSchedule)
		roomScheduleMiddleware.GET("/countScheduleRoom", controller.GetCountScheduleRoom)
		roomScheduleMiddleware.GET("/countScheduleUser", controller.GetcountScheduleUser)
		roomScheduleMiddleware.GET("/:scheduleId", controller.GetScheduleById)
		roomScheduleMiddleware.GET("", controller.GetAllRoomSchedule)
	}
}
