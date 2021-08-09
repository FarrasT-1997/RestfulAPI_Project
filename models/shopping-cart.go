package models

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	Price         int `gorm:"not null"`
	Quantity      int `gorm:"not null"`
	ProductID     uint
	ProductName   string
	TransactionID uint
}
