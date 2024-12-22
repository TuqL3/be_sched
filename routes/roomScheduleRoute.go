package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func RoomScheduleRoute(route *gin.Engine, controller *controllers.RoomScheduleController) {
	roomScheduleMiddleware := route.Group("/api/v1/schedule")
	{
		roomScheduleMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"createSchedule"}), controller.CreateRoomSchedule)
		roomScheduleMiddleware.PUT("/update/:roomScheduleId", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"modifySchedule"}), controller.UpdateRoomSchedule)
		roomScheduleMiddleware.DELETE("/delete/:roomScheduleId", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"deleteSchedule"}), controller.DeleteRoomSchedule)
		roomScheduleMiddleware.GET("/countScheduleRoom", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"viewSchedule"}), controller.GetCountScheduleRoom)
		roomScheduleMiddleware.GET("/countScheduleUser", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"viewSchedule"}), controller.GetcountScheduleUser)
		roomScheduleMiddleware.GET("/:scheduleId", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"viewSchedule"}), controller.GetScheduleById)
		roomScheduleMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"viewSchedule"}), controller.GetAllRoomSchedule)
		roomScheduleMiddleware.POST("/import", controller.ImportScheduleFromExcel)

	}
}
