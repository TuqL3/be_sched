package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func ReportRoute(route *gin.Engine, controller *controllers.ReportController) {
	reportRouteMiddleware := route.Group("/api/v1/report")
	{
		reportRouteMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin", "giangvien"}, []string{"create"}), controller.CreateReport)
		reportRouteMiddleware.PUT("/update/:reportId", controller.UpdateReport)
		reportRouteMiddleware.DELETE("/delete/:reportId", controller.DeleteReport)
		reportRouteMiddleware.GET("/:reportId", controller.GetReportById)
		reportRouteMiddleware.GET("/getCountReportOfRoom", controller.GetCountReportOfRoom)
		reportRouteMiddleware.GET("/getCountReport", controller.GetCountReport)
		reportRouteMiddleware.GET("", controller.GetAllReport)
	}
}
