package controller

import (
	"net/http"
	"yourmodule/database"
	"yourmodule/model"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []model.Product
	if err := database.DB.Preload("ProductCategory").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProduct(c *gin.Context) {
	var product model.Product
	productID := c.Param("id")
	if err := database.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	var product model.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	var product model.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	database.DB.Delete(&product)
	c.JSON(http.StatusNoContent, nil)
}
