package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionStatus string
	TotalQuantity     int
	TotalPrice        int

	//one to many
	//every Transaction has many payment method, and a payment method can only work for one transaction
	PaymentMethods PaymentMethod
	//every transaction has many users, but a user can only have one transaction at a time
	Users User

	//many2many
	ShoppingCarts []ShoppingCart `gorm:"many2many:transaction_shoppingcart;"`
}
