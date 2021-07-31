package models

type ShoppingCart struct {
	ID            uint
	Price         int `gorm:"not null"`
	Quantity      int `gorm:"not null"`
	TransactionID uint
}
