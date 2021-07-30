package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName      string `json:"name" form:"name" gorm:"size:255;not null"`
	Gender        string `json:"gender" form:"gender" gorm:"not null"`
	Username      string `json:"username" form:"username" gorm:"size:100;not null"`
	Email         string `json:"email" form:"email" gorm:"not null"`
	Password      string `json:"password" form:"password" gorm:"not null"`
	Address       string `json:"address" form:"address" gorm:"not null"`
	Token         string `json:"token" form:"token"`
	TransactionID uint
}
