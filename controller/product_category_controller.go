package controller

import (
	"net/http"
	"yourmodule/database"
	"yourmodule/model"

	"github.com/gin-gonic/gin"
)

func GetProductCategories(c *gin.Context) {
	var categories []model.ProductCategory
	if err := database.DB.Preload("Products").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetProductCategory(c *gin.Context) {
	var category model.ProductCategory
	categoryID := c.Param("id")
	if err := database.DB.First(&category, categoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func CreateProductCategory(c *gin.Context) {
	var category model.ProductCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&category)
	c.JSON(http.StatusCreated, category)
}

func UpdateProductCategory(c *gin.Context) {
	categoryID := c.Param("id")
	var category model.ProductCategory
	if err := database.DB.First(&category, categoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product category not found"})
		return
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&category)
	c.JSON(http.StatusOK, category)
}

func DeleteProductCategory(c *gin.Context) {
	categoryID := c.Param("id")
	var category model.ProductCategory
	if err := database.DB.First(&category, categoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product category not found"})
		return
	}
	database.DB.Delete(&category)
	c.JSON(http.StatusNoContent, nil)
}
