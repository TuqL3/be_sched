package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func ReportRoute(route *gin.Engine, controller *controllers.ReportController) {
	reportRouteMiddleware := route.Group("/api/v1/report")
	{
		reportRouteMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban", "giamdoc"}, []string{"createReport"}), controller.CreateReport)
		reportRouteMiddleware.PUT("/update/:reportId", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban", "giamdoc"}, []string{"modifyReport"}), controller.UpdateReport)
		reportRouteMiddleware.DELETE("/delete/:reportId", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban", "giamdoc"}, []string{"deleteReport"}), controller.DeleteReport)
		reportRouteMiddleware.GET("/:reportId", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban", "giamdoc"}, []string{"viewReport"}), controller.GetReportById)
		reportRouteMiddleware.GET("/getCountReportOfRoom", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban", "giamdoc"}, []string{"viewReport"}), controller.GetCountReportOfRoom)
		reportRouteMiddleware.GET("/getCountReport", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban", "giamdoc"}, []string{"viewReport"}), controller.GetCountReport)
		reportRouteMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin", "giangvien", "trucban", "giamdoc"}, []string{"viewReport"}), controller.GetAllReport)
	}
}
