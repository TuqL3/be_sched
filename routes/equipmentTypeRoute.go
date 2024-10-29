package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func EquipmentTypeRoute(route *gin.Engine, controller *controllers.EquipmentTypeController) {
	equipmenttypeMiddleware := route.Group("/api/v1/equipmenttype")
	{

		equipmenttypeMiddleware.GET("/:equipmenttypeId", controller.GetEquipmentTypeById)
		equipmenttypeMiddleware.GET("", controller.GetAllEquipmentType)
		equipmenttypeMiddleware.POST("/create", controller.CreateEquipmentType)
		equipmenttypeMiddleware.PUT("/update/:equipmenttypeId", controller.UpdateEquipmentType)
		equipmenttypeMiddleware.DELETE("/delete/:equipmenttypeId", controller.DeleteEquipmentType)
	}
}
