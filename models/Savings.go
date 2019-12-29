package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	u "github.com/sajicode/gobank/utils"
)

type Savings struct {
	gorm.Model
	AccountBalance int       `gorm:"default:100"json:"account_balance"`
	SaveAmount     int       `json:"save_amount"`
	LastSaveDate   time.Time `json:"last_save_date"gorm:"default:CURRENT_TIMESTAMP"`
	SavingsPlanId  uint      `json:"savings_plan_id"`
	UserId         uint      `json:"user_id"`
}

func (savings *Savings) Validate() (map[string]interface{}, bool) {
	if savings.SaveAmount <= 0 {
		return u.Message(false, "Save Amount must be on the payload"), false
	}

	if savings.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	if savings.SavingsPlanId <= 0 {
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

func GetSaving(id uint) *Savings {
	savings := &Savings{}
	err := GetDB().Table("savings").Where("id = ?", id).First(savings).Error

	if err != nil {
		return nil
	}

	return savings
}

func GetSavings(user uint) []*Savings {
	savings := make([]*Savings, 0)
	err := GetDB().Table("savings").Joins("inner join savings_plans on savings_plans.id = savings.savings_plan_id").Where("user_id = ?", user).Find(&savings).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return savings
}

func (savings *Savings) TopUpSave(savingsId uint) (map[string]interface{}, bool) {

	if resp, ok := savings.Validate(); !ok {
		return resp, true
	}

	err := GetDB().Table("savings").Where("id = ?", savingsId).Updates(Savings{SaveAmount: savings.SaveAmount, AccountBalance: savings.AccountBalance + savings.SaveAmount, LastSaveDate: time.Now()}).Error

	if err != nil {
		fmt.Println(err)
		//* add logger
		resp := u.Message(false, "Database Error")
		return resp, true

	}

	resp := u.Message(true, "success")
	resp["savings"] = savings
	return resp, false

}
