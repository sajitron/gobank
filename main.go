package main

import (
	"fmt"
	"net/http"
	"os"

	"./app"
	"./controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //* attach JWT middleware

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/test", controllers.Test).Methods("GET")

	// router.HandleFunc("/api/contact/new", controllers.CreateContact).Methods("POST")

	// router.HandleFunc("/api/contact/{id}", controllers.GetOneContact).Methods("GET")

	// router.HandleFunc("/api/contacts/{user_id}", controllers.GetContactsFor).Methods("GET")

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
