package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func ReportRoute(route *gin.Engine, controller *controllers.ReportController) {
	reportRouteMiddleware := route.Group("/api/v1/report")
	{

		reportRouteMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin", "giangvien"}, []string{"create"}), controller.CreateReport)
		reportRouteMiddleware.PUT("/update/:reportId", controller.UpdateReport)
		reportRouteMiddleware.DELETE("/delete/:reportId", controller.DeleteReport)
		reportRouteMiddleware.GET("/:reportId", controller.GetReportById)
		reportRouteMiddleware.GET("/getCountReportOfRoom", controller.GetCountReportOfRoom)
		reportRouteMiddleware.GET("", controller.GetAllReport)
	}
}
