package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/suresh-02/Iragu-booking/database"
	"github.com/suresh-02/Iragu-booking/handlers"
	"github.com/suresh-02/Iragu-booking/middleware"
)

func main() {
	// mux router instance for aplication
	router := mux.NewRouter()

	// function to connect to the database
	database.Connect()
	// function to auto migrate and sync schema changes with the database
	database.SyncDatabase()

	// routes
	// Initialize the router

	// Define routes
	router.HandleFunc("/register", handlers.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", handlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/validate", middleware.Validate).Methods(http.MethodPost)

	// Start the server
	log.Println("Server running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", router))

}
