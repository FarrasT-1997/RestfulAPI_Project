package models

type Category struct {
	ID          uint
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Products    []Product
}
