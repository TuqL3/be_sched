package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func RoleRoute(route *gin.Engine, controller *controllers.RoleController) {
	roleRouteMiddleware := route.Group("/api/v1/role")
	{

		roleRouteMiddleware.POST("/create", controller.CreateRole)
		roleRouteMiddleware.DELETE("/delete/:roleId", controller.DeleteRole)
		roleRouteMiddleware.PUT("/update/:roleId", controller.UpdateRole)
		roleRouteMiddleware.GET("/:roleId", controller.GetRoleById)
		roleRouteMiddleware.GET("", controller.GetAllRole)
	}
}
