package database

import (
	model "yourmodule/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	var err error

	// Define the MySQL connection string
	dsn := "root:@tcp(localhost:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the MySQL database
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Auto-migrate the schema (create tables if they don't exist)
	db.AutoMigrate(&model.User{}, &model.Product{}) // Replace User with your model struct
}
