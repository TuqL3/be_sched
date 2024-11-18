package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func RoomScheduleRoute(route *gin.Engine, controller *controllers.RoomScheduleController) {
	roomScheduleMiddleware := route.Group("/api/v1/schedule")
	{
		roomScheduleMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban"}, []string{"createSchedule"}), controller.CreateRoomSchedule)
		roomScheduleMiddleware.PUT("/update/:roomScheduleId", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban"}, []string{"modifySchedule"}), controller.UpdateRoomSchedule)
		roomScheduleMiddleware.DELETE("/delete/:roomScheduleId", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban"}, []string{"deleteSchedule"}), controller.DeleteRoomSchedule)
		roomScheduleMiddleware.GET("/countScheduleRoom", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban"}, []string{"viewSchedule"}), controller.GetCountScheduleRoom)
		roomScheduleMiddleware.GET("/countScheduleUser", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban"}, []string{"viewSchedule"}), controller.GetcountScheduleUser)
		roomScheduleMiddleware.GET("/:scheduleId", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban"}, []string{"viewSchedule"}), controller.GetScheduleById)
		roomScheduleMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban"}, []string{"viewSchedule"}), controller.GetAllRoomSchedule)
	}
}
