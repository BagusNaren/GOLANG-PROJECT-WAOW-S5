package controllers

import (
	"my_project/config"
	"my_project/models"
	"my_project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole.(string) != "USER" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Title      string `json:"title" binding:"required"`
		Content    string `json:"content" binding:"required"`
		CategoryID string `json:"category_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		ArticleID:  utils.GenerateUUID(),
		Title:      input.Title,
		Content:    input.Content,
		CategoryID: input.CategoryID,
		Slug:       utils.GenerateSlug(input.Title),
	}

	if err := config.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	if err := config.DB.Preload("Category").Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve articles: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}