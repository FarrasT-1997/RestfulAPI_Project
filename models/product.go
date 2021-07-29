package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name                 string `gorm:"size:100;not null"`
	CategoryID           uint
	ProductDescriptionID string //getting the ProductDescriptions instead if ID
	Price                int    `gorm:"not null"`
	Stock                int    `gorm:"not null"`

	//belong to
	Category           Category
	ProductDescription ProductDescription `gorm:"references:ProductDescriptions"`

	//many2many
	ShoppingCarts []*ShoppingCart `gorm:"many2many:product_shoppingcart;"`
}
