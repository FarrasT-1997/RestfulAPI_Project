package models

type Product struct {
	ID         uint
	Name       string `gorm:"size:100;not null"`
	CategoryID uint
	Price      int `gorm:"not null"`
	Stock      int `gorm:"not null"`

	//belong to
	ProductDescription ProductDescription

	//many2many
	// ShoppingCarts []*ShoppingCart `gorm:"many2many:product_shoppingcart;"`
}
