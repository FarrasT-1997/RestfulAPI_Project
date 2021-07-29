package models

import (
	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model
	// Transaction   Transaction
	// TransactionId uint `gorm:"not null"`

	// Product   Product
	// ProductId uint `gorm:"not null"`

	// ProductName string `gorm:"not null"`
	Price       int    `gorm:"not null"`
	Quantity    int    `gorm:"not null"`

	Status string `gorm:"not null"`
}
