package models

type ShoppingCart struct {
	ID            uint `gorm:"primaryKey"`
	Price         int  `gorm:"not null"`
	Quantity      int  `gorm:"not null"`
	ProductID     uint
	ProductName   string
	TransactionID uint
}
