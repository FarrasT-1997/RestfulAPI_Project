package models

type Product struct {
	ID                 uint
	Name               string `gorm:"size:100;not null"`
	CategoryID         uint
	Price              int `gorm:"not null"`
	Stock              int `gorm:"not null"`
	ProductDescription ProductDescription
}
