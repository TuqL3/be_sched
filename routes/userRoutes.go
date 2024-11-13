package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func UserRoute(route *gin.Engine, controller *controllers.UserController) {
	publicRoute := route.Group("/api/v1/user")
	{
		publicRoute.POST("/login", controller.Login)
		publicRoute.POST("/register", controller.Register)
	}

	userRouteMiddleware := route.Group("/api/v1/user")
	{

		userRouteMiddleware.PUT("/update/:userId", controller.UpdateUser)
		userRouteMiddleware.DELETE("/delete/:userId", controller.DeleteUser)
		userRouteMiddleware.GET("/getcountuser", controller.GetCountUser)
		userRouteMiddleware.GET("/profile", middleware.RolePermissionMiddleware([]string{"admin", "giangvien"}, []string{"view", "update"}), controller.GetAllUsers)
		userRouteMiddleware.GET("/:userId", controller.GetUserById)
		userRouteMiddleware.GET("", controller.GetAllUsers)
	}
}
