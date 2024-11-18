package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func EquipmentRoute(route *gin.Engine, controller *controllers.EquipmentController) {
	equipmentMiddleware := route.Group("/api/v1/equipment")
	{

		equipmentMiddleware.GET("/equipmentstatus", middleware.RolePermissionMiddleware([]string{"admin", "trucban", "giamdoc"}, []string{"viewEquipment"}), controller.GetQuantityByStatus)
		equipmentMiddleware.GET("/getCountEquipment", middleware.RolePermissionMiddleware([]string{"admin", "trucban", "giamdoc"}, []string{"viewEquipment"}), controller.GetCountEquipment)
		equipmentMiddleware.GET("/:equipmentId", middleware.RolePermissionMiddleware([]string{"admin", "trucban", "giamdoc"}, []string{"viewEquipment"}), controller.GetEquipmentById)
		equipmentMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin", "trucban", "giamdoc"}, []string{"viewEquipment"}), controller.GetAllEquipment)
		equipmentMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin", "trucban"}, []string{"createEquipment"}), controller.CreateEquipment)
		equipmentMiddleware.PUT("/update/:equipmentId", middleware.RolePermissionMiddleware([]string{"admin", "trucban"}, []string{"modifyEquipment"}), controller.UpdateEquipment)
		equipmentMiddleware.DELETE("/delete/:equipmentId", middleware.RolePermissionMiddleware([]string{"admin", "trucban"}, []string{"deleteEquipment"}), controller.DeleteEquipment)
	}
}
