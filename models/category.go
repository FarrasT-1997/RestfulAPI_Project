package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	// Products    []Product `gorm:"one2many:products_categories"`
}
