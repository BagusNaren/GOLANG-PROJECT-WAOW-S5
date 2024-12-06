package main

import (
	"log"
	"os"

	"my_project/config"
	"my_project/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	r := gin.Default()

	routes.UserRoutes(r)
	routes.CategoryRoutes(r)
	routes.ArticleRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}