package models

import (
	"gorm.io/gorm"
)

type ProductDescription struct {
	gorm.Model
	ProductDescriptions string `gorm:"not null"`
}
