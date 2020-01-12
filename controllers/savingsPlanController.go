package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sajicode/gobank/models"
	u "github.com/sajicode/gobank/utils"
)

var CreateSavingsPlan = func(w http.ResponseWriter, r *http.Request) {
	savingsPlan := &models.SavingsPlan{}

	err := json.NewDecoder(r.Body).Decode(savingsPlan)
	if err != nil {
		u.Respond(w, u.Message(false, "Error decoding request body"))
		return
	}

	resp, error := savingsPlan.Create()

	if error == true {
		standardLogger.InvalidRequest("Invalid Request Body to create Savings Plan")
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, resp)
		return
	}

	u.Respond(w, resp)
}

var GetSavingsPlan = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	data, err := models.GetSavingsPlan(id)

	if err == true {
		standardLogger.InvalidRequest("Invalid Request to get savings plan")
		w.WriteHeader(http.StatusBadRequest)
		resp := u.Message(false, "Error getting savings plan")
		u.Respond(w, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetAllSavingsPlans = func(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllSavingsPlans()

	if err == true {
		standardLogger.InvalidRequest("Invalid Request to get savings plans")
		w.WriteHeader(http.StatusBadRequest)
		resp := u.Message(false, "Error getting savings plans")
		u.Respond(w, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
