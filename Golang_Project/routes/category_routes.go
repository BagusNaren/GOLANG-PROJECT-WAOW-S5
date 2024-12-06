package routes

import (
	"my_project/controllers"
	"my_project/middleware"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	category := r.Group("/categories")
	{
		category.Use(middleware.AuthMiddleware)
		category.POST("/", controllers.CreateCategory)
		category.GET("/", controllers.GetAllCategories)
	}
}