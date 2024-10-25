package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func ReportRoute(route *gin.Engine, controller *controllers.ReportController) {
	reportRouteMiddleware := route.Group("/api/v1/report")
	{
		reportRouteMiddleware.Use(middleware.AuthMiddleware())

		reportRouteMiddleware.POST("/create", controller.CreateReport)
		reportRouteMiddleware.PUT("/update/:reportId", controller.UpdateReport)
		reportRouteMiddleware.DELETE("/delete/:reportId", controller.DeleteReport)
		reportRouteMiddleware.GET("", controller.GetAllReport)
	}
}
