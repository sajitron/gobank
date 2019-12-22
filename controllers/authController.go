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

	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

var Test = func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "success")
	resp["data"] = "Test Passed"
	u.Respond(w, resp)
}
