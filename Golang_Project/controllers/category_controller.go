package controllers

import (
	"my_project/models"
	"my_project/config"
	"my_project/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCategory(c *gin.Context) {
	userRole := c.MustGet("role").(string)
	if userRole != "USER" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		CategoryID:  utils.GenerateUUID(),
		Name:        input.Name,
		Slug:        utils.GenerateSlug(input.Name),
		Description: input.Description,
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}