package models

type PaymentMethodTable struct {
	ID   uint
	Name string `gorm:"size:255;not null"`
}
