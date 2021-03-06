package database

import (
	"restfulAPI/config"
	"restfulAPI/models"
)

func CreateTransaction(transaction models.Transaction) (interface{}, error) {

	if err := config.DB.Save(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func GetTransactionById(name string) (interface{}, error) {
	var transaction []models.Transaction
	if err := config.DB.Find(&transaction, "users=?", name).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func InputTransactioData(userList models.User, payment []models.PaymentMethodTable) models.Transaction {
	transaction := models.Transaction{
		Users:             userList.FullName,
		Address:           userList.Address,
		PaymentMethodID:   99,
		PaymentMethod:     payment[len(payment)-1],
		TransactionStatus: "waiting",
		TotalQuantity:     0,
		TotalPrice:        0,
		Checkout:          "not yet",
	}

	return transaction
}

func DeleteTransactions(id int) (interface{}, error) {
	var deleteTransaction models.Transaction
	if err := config.DB.Find(&deleteTransaction, "id=?", id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&deleteTransaction).Error; err != nil {
		return nil, err
	}
	return deleteTransaction, nil
}

func GetOneTransaction(id int) (models.Transaction, error) {
	var transaction models.Transaction
	if err := config.DB.Find(&transaction, "id=?", id).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func EditTransaction(transaction models.Transaction) (interface{}, error) {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
