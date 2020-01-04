package models

import (
	"fmt"
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

func GetTransaction(id string) *Transaction {
	transaction := &Transaction{}

	//* add logger
	err := GetDB().Table("transaction").Where("id = ?", id).First(transaction).Error

	if err != nil {
		fmt.Printf("model error: %v", err)
		return nil
	}

	return transaction
}

func GetTransactions(savings_id string) []*Transaction {
	transaction := make([]*Transaction, 0)
	err := GetDB().Table("transaction").Joins("inner join savings_plans on savings_plans.id = savings.savings_plan_id").Where("savings_id = ?", savings_id).Find(&transaction).Error

	//* add logger
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return transaction
}
