package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

var CreateSaving = func(w http.ResponseWriter, r *http.Request) {
	savings := &models.Savings{}

	err := json.NewDecoder(r.Body).Decode(savings)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error decoding request body"))
		return
	}

	resp := savings.Create()
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
