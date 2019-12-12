package models

import (
	"fmt"
	"time"

	u "../utils"
	"github.com/jinzhu/gorm"
)

type Savings struct {
	gorm.Model
	AccountBalance *int      `gorm:"default:100"`
	SaveAmount     int       `json:"save_amount"`
	LastSaveDate   time.Time `json:"last_save_date"`
	SavingsPlanId  uint      `json:"savings_plan_id"`
	UserId         uint      `json:"user_id"`
}

func (savings *Savings) Validate() (map[string]interface{}, bool) {
	if savings.SaveAmount <= 0 {
		return u.Message(false, "Save Amount must be on the payload"), false
	}

	if savings.LastSaveDate.IsZero() {
		return u.Message(false, "Last save date must be on the payload"), false
	}

	if savings.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	if savings.SavingsPlanId <= 0 {
		return u.Message(false, "Savings Plan is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (savings *Savings) Create() map[string]interface{} {
	if resp, ok := savings.Validate(); !ok {
		return resp
	}

	GetDB().Create(savings)

	resp := u.Message(true, "success")
	resp["savings"] = savings
	return resp
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
	err := GetDB().Table("savings").Where("user_id = ?", user).Find(&savings).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return savings
}
