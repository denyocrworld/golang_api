package model

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	ProductName string `gorm:"not null"`
	Price       float32
}
