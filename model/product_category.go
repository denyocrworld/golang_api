package model

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	ID           uint `json:"id" gorm:"primaryKey"`
	CategoryName string
	Products     []Product // Relasi satu ke banyak dengan Produk
}
