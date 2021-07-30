package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"name" form:"name" gorm:"size:255;not null"`
	Gender   string `gorm:"size:1;not null"`
	Username string `gorm:"size:100;not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Address  string `gorm:"not null"`

	TransactionID uint
}
