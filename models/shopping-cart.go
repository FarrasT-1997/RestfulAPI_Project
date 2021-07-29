package models

import (
	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model

	Price    int `gorm:"not null"`
	Quantity int `gorm:"not null"`

	Status string `gorm:"not null"`

	Transactions []*Transaction `gorm:"many2many:transaction_shoppingcart;"`
	Products     []*Product     `gorm:"many2many:product_shoppingcart;"`
}
