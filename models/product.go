package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string `gorm:"size:100;not null"`
	// Description string `gorm:"not null"`
	Price int `gorm:"not null"`
	Stock int `gorm:"not null"`

	Categories          []Category           `gorm:"foreignKey:CategoryID"`
	ProductDescriptions []ProductDescription `gorm:"foreignKey:ProductDescriptionID"`
}
