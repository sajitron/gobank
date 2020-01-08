package models

import (
	"fmt"

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

func GetSaving(id string) *Savings {
	savings := &Savings{}

	//* add logger
	err := GetDB().Table("savings").Where("id = ?", id).First(savings).Error

	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		fmt.Printf("model error: %v", err)
		return nil
	}

	return savings
}

func GetSavings(account string) []*Savings {
	savings := make([]*Savings, 0)
	err := GetDB().Table("savings").Joins("inner join savings_plans on savings_plans.id = savings.savings_plan_id").Where("account_id = ?", account).Find(&savings).Error

	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		fmt.Println(err)
		return nil
	}

	return savings
}

func (savings *Savings) TopUpSave(savings_id string, amount int) (map[string]interface{}, bool) {

	if resp, ok := savings.Validate(); !ok {
		return resp, true
	}

	//* create transaction
	transaction := Transaction{
		SavingsId: savings_id,
		Amount:    amount,
	}
	err := GetDB().Create(transaction)

	if err != nil {
		fmt.Println(err)
		standardLogger.InvalidRequest(err.Error.Error())
		resp := u.Message(false, "Database Error")
		return resp, true
	}

	//* update savings balance

	GetDB().Table("savings").Update("account_balance", gorm.Expr("account_balance + ?", amount)).Where("id = ?", savings_id)

	if err != nil {
		fmt.Println(err)
		standardLogger.InvalidRequest(err.Error.Error())
		resp := u.Message(false, "Database Error")
		return resp, true
	}

	resp := u.Message(true, "success")
	resp["savings"] = savings
	return resp, false

}
