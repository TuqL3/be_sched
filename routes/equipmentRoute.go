package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func EquipmentRoute(route *gin.Engine, controller *controllers.EquipmentController) {
	equipmentMiddleware := route.Group("/api/v1/equipment")
	{
		equipmentMiddleware.Use(middleware.AuthMiddleware())
		equipmentMiddleware.Use(middleware.AdminOnly())

		equipmentMiddleware.POST("/create", controller.CreateEquipment)
		equipmentMiddleware.PUT("/update/:equipmentId", controller.UpdateEquipment)
		equipmentMiddleware.DELETE("/delete/:equipmentId", controller.DeleteEquipment)
	}
}
