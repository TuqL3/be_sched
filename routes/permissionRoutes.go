package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func PermissionRoute(route *gin.Engine, controller *controllers.PermissionController) {
	permissionRouteMiddleware := route.Group("/api/v1/permission")
	{

		permissionRouteMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"createPermission"}), controller.CreatePermission)
		permissionRouteMiddleware.DELETE("/delete/:permissionId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"deletePermission"}), controller.DeletePermission)
		permissionRouteMiddleware.PUT("/update/:permissionId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"modifyPermission"}), controller.UpdatePermission)
		permissionRouteMiddleware.GET("/:permissionId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"viewPermission"}), controller.GetPermissionById)
		permissionRouteMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"viewPermission"}), controller.GetAllPermission)
	}
}
