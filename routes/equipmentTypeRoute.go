package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func EquipmentTypeRoute(route *gin.Engine, controller *controllers.EquipmentTypeController) {
	equipmenttypeMiddleware := route.Group("/api/v1/equipmenttype")
	{

		equipmenttypeMiddleware.GET("/:equipmenttypeId", middleware.RolePermissionMiddleware([]string{"admin", "trucban", "giamdoc"}, []string{"viewEquipmentType"}), controller.GetEquipmentTypeById)
		equipmenttypeMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin", "trucban", "giamdoc"}, []string{"viewEquipmentType"}), controller.GetAllEquipmentType)
		equipmenttypeMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin", "trucban", "giamdoc"}, []string{"createEquipmentType"}), controller.CreateEquipmentType)
		equipmenttypeMiddleware.PUT("/update/:equipmenttypeId", middleware.RolePermissionMiddleware([]string{"admin", "trucban"}, []string{"modifyEquipmentType"}), controller.UpdateEquipmentType)
		equipmenttypeMiddleware.DELETE("/delete/:equipmenttypeId", middleware.RolePermissionMiddleware([]string{"admin", "trucban"}, []string{"deleteEquipmentType"}), controller.DeleteEquipmentType)
	}
}
