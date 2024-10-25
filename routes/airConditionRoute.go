package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func AirConditionRoute(route *gin.Engine, controller *controllers.AirConditionController) {
	airConditionMiddleware := route.Group("/api/v1/aircondition")
	{
		airConditionMiddleware.Use(middleware.AuthMiddleware())
		airConditionMiddleware.Use(middleware.AdminOnly())

		airConditionMiddleware.GET("", controller.GetAllAirCondition)
		airConditionMiddleware.POST("/create", controller.CreateAirCondition)
		airConditionMiddleware.PUT("/update/:airConditionId", controller.UpdateAirCondition)
		airConditionMiddleware.DELETE("/delete/:airConditionId", controller.DeleteAirCondition)
	}
}
