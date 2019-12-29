package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "github.com/sajicode/logger"
	"github.com/sajicode/models"
	u "github.com/sajicode/utils"
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
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Request error"))
		return
	}
	data := models.GetSaving(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetAllSavings = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["user_id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Request Error"))
		return
	}

	data := models.GetSavings(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var TopUpSavings = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["savings_id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Request error"))
		return
	}

	savings := &models.Savings{}

	err = json.NewDecoder(r.Body).Decode(savings)

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error decoding request body"))
		return
	}

	resp, error := savings.TopUpSave(uint(id))

	if error == true {
		standardLogger.InvalidRequest("Invalid Request to Top up saving")
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, resp)
		return
	}

	u.Respond(w, resp)
}
