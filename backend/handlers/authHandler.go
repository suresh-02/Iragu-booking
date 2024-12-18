package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/suresh-02/Iragu-booking/database"
	"github.com/suresh-02/Iragu-booking/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	// Define the request body structure
	var signupBody struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&signupBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupBody.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to create hashed password", http.StatusInternalServerError)
		return
	}

	// Create the user object
	signupUser := models.UserCreds{
		Username: signupBody.Username,
		Password: string(hashedPassword),
		Email:    signupBody.Email,
	}

	// Save the user to the database
	result := database.DB.Create(&signupUser)
	if result.Error != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}
