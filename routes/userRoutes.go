package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine, controller *controllers.UserController) {
	publicRoute := route.Group("/api/v1/user")
	{
		publicRoute.POST("/login", controller.Login)
		publicRoute.POST("/register", middleware.RolePermissionMiddleware([]string{"admin"}, []string{"createUser"}), controller.Register)
	}

	userRouteMiddleware := route.Group("/api/v1/user")
	{

		userRouteMiddleware.PUT("/update/:userId", controller.UpdateUser)
		userRouteMiddleware.DELETE("/delete/:userId", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"deleteUser"}), controller.DeleteUser)
		userRouteMiddleware.GET("/getcountuser", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"viewUser"}), controller.GetCountUser)
		userRouteMiddleware.GET("/profile", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien", "giangvien", "trucban", "giamdoc"}, []string{"viewProfile"}), controller.GetAllUsers)
		userRouteMiddleware.GET("/:userId", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"viewUser"}), controller.GetUserById)
		userRouteMiddleware.GET("", middleware.RolePermissionMiddleware([]string{"admin", "giamdoc", "trucban", "giangvien"}, []string{"viewUser"}), controller.GetAllUsers)
		userRouteMiddleware.POST("/import", controller.ImportUserFromExcel)
	}
}
