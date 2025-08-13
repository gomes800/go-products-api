package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    uint    `gorm:"primarykey"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
