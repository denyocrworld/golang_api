package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name              string
	Description       string
	Price             float64
	ProductCategoryID uint // Kunci asing untuk menghubungkan ke kategori produk
}
