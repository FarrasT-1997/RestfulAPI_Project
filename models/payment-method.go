package models

import (
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Name          string `gorm:"size:255;not null"`
	Status        string `gorm:"default:null"`
	TransactionID uint
}
