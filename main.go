package main

import (
	"yourmodule/controller"
	"yourmodule/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database MySQL
	database.Connect()

	// Membuat router Gin
	router := gin.Default()

	// Rute untuk mendapatkan semua pengguna
	router.GET("/users", controller.GetUsers)

	// Rute untuk mendapatkan pengguna berdasarkan ID
	router.GET("/users/:id", controller.GetUser)

	// Rute untuk membuat pengguna baru
	router.POST("/users", controller.CreateUser)

	// Rute untuk memperbarui pengguna berdasarkan ID
	router.PUT("/users/:id", controller.UpdateUser)

	// Rute untuk menghapus pengguna berdasarkan ID
	router.DELETE("/users/:id", controller.DeleteUser)

	// Menjalankan server pada port 8080
	router.Run(":8080")
}
