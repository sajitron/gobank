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
	data := models.GetSaving(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetAllSavings = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	data := models.GetSavings(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// var TopUpSavings = func(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["savings_id"]

// 	savings := &models.Savings{}

// 	err := json.NewDecoder(r.Body).Decode(savings)

// 	if err != nil {
// 		fmt.Println(err)
// 		u.Respond(w, u.Message(false, "Error decoding request body"))
// 		return
// 	}

// 	resp, error := savings.TopUpSave(id)

// 	if error == true {
// 		standardLogger.InvalidRequest("Invalid Request to Top up saving")
// 		w.WriteHeader(http.StatusBadRequest)
// 		u.Respond(w, resp)
// 		return
// 	}

// 	u.Respond(w, resp)
// }
