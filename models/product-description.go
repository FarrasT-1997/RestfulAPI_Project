package models

import (
	"gorm.io/gorm"
)

type ProductDescription struct {
	gorm.Model
	ProductDescription string `gorm:"not null"`
}
