package models

import (
	u "../utils"
	"github.com/jinzhu/gorm"
)

type SavingsPlan struct {
	gorm.Model
	Name      string `json:"name"`
	DuePeriod int    `json:"due_period"`
	SaveRate  int    `json:"save_rate"`
}

func (savingsPlan *SavingsPlan) Validate() (map[string]interface{}, bool) {
	if savingsPlan.Name == "" {
		return u.Message(false, "Savings Plan name must be on the payload"), false
	}

	if savingsPlan.DuePeriod < 0 {
		return u.Message(false, "Due Period must be on the payload"), false
	}

	if savingsPlan.SaveRate < 0 {
		return u.Message(false, "Savings rate must be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (savingsPlan *SavingsPlan) Create() map[string]interface{} {
	if resp, ok := savingsPlan.Validate(); !ok {
		return resp
	}

	GetDB().Create(savingsPlan)

	resp := u.Message(true, "success")
	resp["savings_plan"] = savingsPlan
	return resp
}

func GetSavingsPlan(id uint) *SavingsPlan {
	savingsPlan := &SavingsPlan{}
	err := GetDB().Table("savings_plans").Where("id = ?", id).First(savingsPlan).Error

	if err != nil {
		return nil
	}

	return savingsPlan
}
