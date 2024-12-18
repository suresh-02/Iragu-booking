package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupBody.Password), 10)
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

func Login(w http.ResponseWriter, r *http.Request) {

	var loginBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&loginBody)

	var user models.UserCreds

	database.DB.First(&user, "email=?", loginBody.Email)

	if user.ID == 0 {
		http.Error(w, "No user found", http.StatusBadRequest)
		return
	}

	fmt.Println(user.Password)
	// http.Error(w, user.Password, http.StatusAccepted)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusForbidden)
		return
	}

	//  generate a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(time.Hour).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
    if err != nil {
        http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": tokenString,
    })

}
