package models

import (
	"github.com/jinzhu/gorm"
	u "github.com/sajicode/gobank/utils"
)

type Savings struct {
	Base
	AccountBalance int    `gorm:"default:100"json:"account_balance"`
	SavingsPlanId  string `json:"savings_plan_id"`
	AccountId      string `json:"account_id"`
}

func (savings *Savings) Validate() (map[string]interface{}, bool) {

	if savings.AccountId == "" {
		return u.Message(false, "Account is not recognized"), false
	}

	if savings.SavingsPlanId == "" {
		return u.Message(false, "Savings Plan is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (savings *Savings) Create() (map[string]interface{}, bool) {
	if resp, ok := savings.Validate(); !ok {
		return resp, true
	}

	GetDB().Create(savings)

	resp := u.Message(true, "success")
	resp["savings"] = savings
	return resp, false
}

func GetSaving(id string) (*Savings, bool) {
	savings := &Savings{}

	//* add logger
	err := GetDB().Table("savings").Where("id = ?", id).First(savings).Error

	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		return nil, true
	}

	return savings, false
}

func GetSavings(account string) ([]*Savings, bool) {
	savings := make([]*Savings, 0)
	err := GetDB().Table("savings").Where("account_id = ?", account).Find(&savings).Error

	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		return nil, true
	}

	return savings, false
}

func (savings *Savings) TopUpSave(savings_id string, amount int) (map[string]interface{}, bool) {

	if savings_id == "" {
		standardLogger.InvalidRequest("No savings selected")
		resp := u.Message(false, "Incomplete request. No savings selected")
		return resp, true
	}

	if amount <= 0 {
		standardLogger.InvalidRequest("Invalid amount")
		resp := u.Message(false, "Incomplete request. Enter a valid amount")
		return resp, true
	}

	//* create transaction
	transaction := Transaction{
		SavingsId: savings_id,
		Amount:    amount,
	}

	resp, err := transaction.Create()

	if err == true {
		standardLogger.InvalidRequest("Invalid Request Body to Save")
	}

	//* update savings balance

	GetDB().Table("savings").Update("account_balance", gorm.Expr("account_balance + ?", amount)).Where("id = ?", savings_id)

	resp = u.Message(true, "success")
	resp["savings"] = transaction
	return resp, false

}
