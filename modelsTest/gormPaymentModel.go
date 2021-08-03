package modelsTest

import "gorm.io/gorm"

type GormPaymentModel struct {
	db *gorm.DB
}

func (m *GormPaymentModel) Get() []Payment {
	var payments []Payment
	m.db.Find(&payments)
	return payments
}

func (m *GormPaymentModel) Insert(payments Payment) error {
	tx := m.db.Save(&payments)
	return tx.Error
}

func NewGormPaymentModel(db *gorm.DB) *GormPaymentModel {
	return &GormPaymentModel{db: db}
}
