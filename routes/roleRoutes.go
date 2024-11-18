package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func RoleRoute(route *gin.Engine, controller *controllers.RoleController) {
	roleRouteMiddleware := route.Group("/api/v1/role")
	{

		roleRouteMiddleware.POST("/create", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"createRole"}), controller.CreateRole)
		roleRouteMiddleware.DELETE("/delete/:roleId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"deleteRole"}), controller.DeleteRole)
		roleRouteMiddleware.PUT("/update/:roleId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"modifyRole"}), controller.UpdateRole)
		roleRouteMiddleware.GET("/:roleId", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"viewRole"}), controller.GetRoleById)
		roleRouteMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"viewRole"}), controller.GetAllRole)
	}
}
