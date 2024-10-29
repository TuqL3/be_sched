package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func PermissionRoute(route *gin.Engine, controller *controllers.PermissionController) {
	permissionRouteMiddleware := route.Group("/api/v1/permission")
	{

		permissionRouteMiddleware.POST("/create", controller.CreatePermission)
		permissionRouteMiddleware.DELETE("/delete/:permissionId", controller.DeletePermission)
		permissionRouteMiddleware.PUT("/update/:permissionId", controller.UpdatePermission)
		permissionRouteMiddleware.GET("/:permissionId", controller.GetPermissionById)
		permissionRouteMiddleware.GET("", controller.GetAllPermission)
	}
}
