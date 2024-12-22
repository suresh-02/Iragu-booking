package main

import (
	"log"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
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

	corsHandler := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// routes
	// Initialize the router

	// Define routes
	router.HandleFunc("/register", handlers.Register).Methods(http.MethodPost
	router.Handle("/login", middleware.Validate(http.HandlerFunc(handlers.Login))).Methods(http.MethodPost)

	// Start the server
	log.Println("Server running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", corsHandler(router)))

}
