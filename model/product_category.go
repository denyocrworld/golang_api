package model

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name     string
	Products []Product // Relasi satu ke banyak dengan Produk
}
