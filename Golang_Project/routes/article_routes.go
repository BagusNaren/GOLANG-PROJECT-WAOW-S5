package routes

import (
	"my_project/controllers"
	"my_project/middleware"
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.Engine) {
	article := r.Group("/articles")
	{
		article.Use(middleware.AuthMiddleware)
		article.POST("/", controllers.CreateArticle)
		article.GET("/", controllers.GetAllArticles)
	}
}