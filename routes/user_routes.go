package routes

import (
	"my_project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("/register", controllers.Register)
	}
}