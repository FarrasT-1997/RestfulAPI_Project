package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionStatus string
	TotalQuantity     int
	TotalPrice        int
	PaymentMethod     int `gorm:"foreignKey:PaymentMethodID"`
	UserID            int `gorm:"foreignKey:UserID"`
}
