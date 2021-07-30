package models

type ProductDescription struct {
	ID                  uint
	ProductDescriptions string `gorm:"not null"`
	ProductID           uint
}
