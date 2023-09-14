package main

import (
	"net/http"
	"yourmodule/database"
	"yourmodule/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database MySQL
	database.Connect()

	// Membuat router Gin
	router := gin.Default()

	// Rute untuk mendapatkan semua pengguna
	router.GET("/users", getUsers)

	// Rute untuk mendapatkan pengguna berdasarkan ID
	router.GET("/users/:id", getUser)

	// Rute untuk membuat pengguna baru
	router.POST("/users", createUser)

	// Rute untuk memperbarui pengguna berdasarkan ID
	router.PUT("/users/:id", updateUser)

	// Rute untuk menghapus pengguna berdasarkan ID
	router.DELETE("/users/:id", deleteUser)

	// Menjalankan server pada port 8080
	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Handler untuk mendapatkan pengguna berdasarkan ID
func getUser(c *gin.Context) {
	var user model.User
	userID := c.Param("id")
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Handler untuk membuat pengguna baru
func createUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// Handler untuk memperbarui pengguna berdasarkan ID
func updateUser(c *gin.Context) {
	userID := c.Param("id")
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// Handler untuk menghapus pengguna berdasarkan ID
func deleteUser(c *gin.Context) {
	userID := c.Param("id")
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusNoContent, nil)
}
