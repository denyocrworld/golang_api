package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID                uint `json:"id" gorm:"primaryKey"`
	ProductName       string
	Description       string
	Price             float64
	ProductCategory   ProductCategory // Relasi dengan ProductCategory
	ProductCategoryID uint            // Kunci asing untuk menghubungkan ke kategori produk
}
