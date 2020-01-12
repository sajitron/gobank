package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/sajicode/logger"
	"github.com/sajicode/gobank/models"
	u "github.com/sajicode/gobank/utils"
)

//* logger
// var standardLogger = logger.NewLogger()

var CreateSaving = func(w http.ResponseWriter, r *http.Request) {
	savings := &models.Savings{}

	err := json.NewDecoder(r.Body).Decode(savings)
	if err != nil {
		standardLogger.InvalidRequest(err.Error())
		u.Respond(w, u.Message(false, "Error decoding request body"))
		return
	}

	resp, error := savings.Create()

	if error == true {
		standardLogger.InvalidRequest("Invalid Request Body to Save")
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, resp)
		return
	}

	u.Respond(w, resp)
}

var GetSaving = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, err := models.GetSaving(id)

	if err == true {
		standardLogger.InvalidRequest("Invalid Request to get savings")
		w.WriteHeader(http.StatusBadRequest)
		resp := u.Message(false, "Error getting savings")
		u.Respond(w, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

//* get all of a users savings
var GetAllSavings = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	data, err := models.GetSavings(id)

	if err == true {
		standardLogger.InvalidRequest("Invalid Request to get savings")
		w.WriteHeader(http.StatusBadRequest)
		resp := u.Message(false, "Error getting savings")
		u.Respond(w, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
