package models

import (
	u "github.com/sajicode/gobank/utils"
)

type SavingsPlan struct {
	Base
	Name      string `json:"name"`
	DuePeriod int    `json:"due_period"`
	SaveRate  int    `json:"save_rate"`
}

func (savingsPlan *SavingsPlan) Validate() (map[string]interface{}, bool) {
	if savingsPlan.Name == "" {
		return u.Message(false, "Savings Plan name must be on the payload"), false
	}

	if savingsPlan.DuePeriod <= 0 {
		return u.Message(false, "Due Period must be on the payload"), false
	}

	if savingsPlan.SaveRate <= 0 {
		return u.Message(false, "Savings rate must be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (savingsPlan *SavingsPlan) Create() (map[string]interface{}, bool) {
	if resp, ok := savingsPlan.Validate(); !ok {
		return resp, true
	}

	GetDB().Create(savingsPlan)

	resp := u.Message(true, "success")
	resp["savings_plan"] = savingsPlan
	return resp, false
}

func GetSavingsPlan(id string) (*SavingsPlan, bool) {
	savingsPlan := &SavingsPlan{}
	err := GetDB().Table("savings_plans").Where("id = ?", id).First(savingsPlan).Error

	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		return nil, true
	}

	return savingsPlan, false
}

func GetAllSavingsPlans() ([]*SavingsPlan, bool) {
	savingsPlans := make([]*SavingsPlan, 0)
	err := GetDB().Table("savings_plans").Find(&savingsPlans).Error

	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		return nil, true
	}

	return savingsPlans, false
}
