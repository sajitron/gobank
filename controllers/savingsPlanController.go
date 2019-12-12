package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

var CreateSavingsPlan = func(w http.ResponseWriter, r *http.Request) {
	savingsPlan := &models.SavingsPlan{}

	err := json.NewDecoder(r.Body).Decode(savingsPlan)
	if err != nil {
		u.Respond(w, u.Message(false, "Error decoding request body"))
		return
	}

	resp := savingsPlan.Create()
	u.Respond(w, resp)
}

var GetSavingsPlan = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Request error"))
		return
	}
	data := models.GetSavingsPlan(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetAllSavingsPlans = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllSavingsPlans()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
