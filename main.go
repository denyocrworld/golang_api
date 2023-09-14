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

	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUser)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)

	router.GET("/products", controller.GetProducts)
	router.GET("/products/:id", controller.GetProduct)
	router.POST("/products", controller.CreateProduct)
	router.PUT("/products/:id", controller.UpdateProduct)
	router.DELETE("/products/:id", controller.DeleteProduct)

	router.GET("/product_categories", controller.GetProductCategories)
	router.GET("/product_categories/:id", controller.GetProductCategory)
	router.POST("/product_categories", controller.CreateProductCategory)
	router.PUT("/product_categories/:id", controller.UpdateProductCategory)
	router.DELETE("/product_categories/:id", controller.DeleteProductCategory)

	// Menjalankan server pada port 8080
	router.Run(":8080")
}
