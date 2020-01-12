package models

import (
	u "github.com/sajicode/gobank/utils"
)

type Transaction struct {
	Base
	Amount    int    `json:"amount"`
	SavingsId string `json:"savings_id"`
}

func (transaction *Transaction) Validate() (map[string]interface{}, bool) {
	if transaction.Amount <= 0 {
		return u.Message(false, "Save Amount must be on the payload"), false
	}

	if transaction.SavingsId == "" {
		return u.Message(false, "Savings not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (transaction *Transaction) Create() (map[string]interface{}, bool) {
	if resp, ok := transaction.Validate(); !ok {
		return resp, true
	}

	GetDB().Create(transaction)

	resp := u.Message(true, "success")
	resp["transaction"] = transaction
	return resp, false
}

func GetTransaction(id string) (*Transaction, bool) {
	transaction := &Transaction{}

	//* add logger
	err := GetDB().Table("transactions").Where("id = ?", id).First(transaction).Error

	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		return nil, true
	}

	return transaction, false
}

func GetTransactions(savings_id string) ([]*Transaction, bool) {
	transaction := make([]*Transaction, 0)
	err := GetDB().Table("transactions").Where("savings_id = ?", savings_id).Find(&transaction).Error

	//* add logger
	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		return nil, true
	}

	return transaction, false
}
