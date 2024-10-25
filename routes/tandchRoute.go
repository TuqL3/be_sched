package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func TAndChtRoute(route *gin.Engine, controller *controllers.TAndChController) {
	tandchMiddleware := route.Group("/api/v1/tandch")
	{
		tandchMiddleware.Use(middleware.AuthMiddleware())
		tandchMiddleware.Use(middleware.AdminOnly())

		tandchMiddleware.POST("/create", controller.CreateTAndCh)
		tandchMiddleware.PUT("/update/:tAndChId", controller.UpdateTAndCh)
		tandchMiddleware.DELETE("/delete/:tAndChId", controller.DeleteTAndCh)
		tandchMiddleware.GET("", controller.GetAllTAndCh)
	}
}
