package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func EquipmentRoute(route *gin.Engine, controller *controllers.EquipmentController) {
	equipmentMiddleware := route.Group("/api/v1/equipment")
	{

		equipmentMiddleware.GET("/equipmentstatus", controller.GetQuantityByStatus)
		equipmentMiddleware.GET("/:equipmentId", controller.GetEquipmentById)
		equipmentMiddleware.GET("", controller.GetAllEquipment)
		equipmentMiddleware.POST("/create", controller.CreateEquipment)
		equipmentMiddleware.PUT("/update/:equipmentId", controller.UpdateEquipment)
		equipmentMiddleware.DELETE("/delete/:equipmentId", controller.DeleteEquipment)
	}
}
