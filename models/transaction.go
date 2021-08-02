package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionStatus string
	TotalQuantity     int
	TotalPrice        int
	PaymentMethodID   uint
	PaymentMethod     PaymentMethodTable
	Users             string
	Address           string
}
