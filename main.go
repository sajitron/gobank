package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sajicode/app"
	"github.com/sajicode/controllers"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //* attach JWT middleware

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/test", controllers.Test).Methods("GET")

	router.HandleFunc("/api/plan", controllers.CreateSavingsPlan).Methods("POST")

	router.HandleFunc("/api/plan/{id}", controllers.GetSavingsPlan).Methods("GET")

	router.HandleFunc("/api/plan", controllers.GetAllSavingsPlans).Methods("GET")

	router.HandleFunc("/api/save", controllers.CreateSaving).Methods("POST")

	router.HandleFunc("/api/topup/{savings_id}", controllers.TopUpSavings).Methods("PUT")

	router.HandleFunc("/api/save/{id}", controllers.GetSaving).Methods("GET")

	router.HandleFunc("/api/saves/{user_id}", controllers.GetAllSavings).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
