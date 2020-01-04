package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/sajicode/logger"
	"github.com/sajicode/gobank/models"
	u "github.com/sajicode/gobank/utils"
)

var TopUpSavings = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["savings_id"]

	savings := &models.Savings{}
	transaction := &models.Transaction{}

	err := json.NewDecoder(r.Body).Decode(transaction)

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error decoding request body"))
		return
	}

	resp, error := savings.TopUpSave(id, transaction.Amount)

	if error == true {
		standardLogger.InvalidRequest("Invalid Request to Top up saving")
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, resp)
		return
	}

	u.Respond(w, resp)

}
