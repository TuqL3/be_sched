package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func CategoryRoute(route *gin.Engine, controller *controllers.CategoryController) {
	categoryMiddleware := route.Group("/api/v1/category")
	{
		categoryMiddleware.Use(middleware.AuthMiddleware())
		categoryMiddleware.Use(middleware.AdminOnly())

		categoryMiddleware.GET("/:categoryId", controller.GetCategoryById)
		categoryMiddleware.GET("", controller.GetAllCategory)
		categoryMiddleware.POST("/create", controller.CreateCategory)
		categoryMiddleware.PUT("/update/:categoryId", controller.UpdateCategory)
		categoryMiddleware.DELETE("/delete/:categoryId", controller.DeleteCategory)
	}
}
