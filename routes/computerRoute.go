package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func ComputerRoute(route *gin.Engine, controller *controllers.ComputerController) {
	computerMiddleware := route.Group("/api/v1/computer")
	{
		computerMiddleware.Use(middleware.AuthMiddleware())
		computerMiddleware.Use(middleware.AdminOnly())

		computerMiddleware.POST("/create", controller.CreateComputer)
		computerMiddleware.PUT("/update/:computerId", controller.UpdateComputer)
		computerMiddleware.DELETE("/delete/:computerId", controller.DeleteComputer)
		computerMiddleware.GET("/", controller.GetAllComputer)
	}
}
