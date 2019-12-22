package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sajicode/models"
	u "github.com/sajicode/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //* decode the request body into struct
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp, error := account.Create()
	if error == true {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, resp)
		return
	}

	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp, error := models.Login(account.Email, account.Password)
		if error == true {
		w.WriteHeader(http.StatusForbidden)
		u.Respond(w, resp)
		return
	}
	u.Respond(w, resp)
}

var Test = func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "success")
	resp["data"] = "Test Passed"
	u.Respond(w, resp)
}
