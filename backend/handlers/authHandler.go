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

func Register(w http.ResponseWriter, r *http.Request) {
	// Define the request body structure
	var reqBody struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
	if err != nil {
		http.Error(w, "Invalid Email or Password", http.StatusBadRequest)
		return
	}

	// Create the user object
	newUser := models.UserCreds{
		Username: reqBody.Username,
		Password: string(hashedPassword),
		Email:    reqBody.Email,
	}

	// Save the user to the database
	result := database.DB.Create(&newUser)
	if result.Error != nil {
		http.Error(w, "Invalid Email or Password", http.StatusBadRequest)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	var reqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&reqBody)

	var user models.UserCreds

	database.DB.First(&user, "email=?", reqBody.Email)

	if user.ID == 0 {
		http.Error(w, "Invalid Email or Password", http.StatusBadRequest)
		return
	}

	fmt.Println(user.Password)
	// http.Error(w, user.Password, http.StatusAccepted)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
	if err != nil {
		http.Error(w, "Invalid Email or Password", http.StatusBadRequest)
		return
	}

	//  generate a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		http.Error(w, "Invalid Email or Password", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		HttpOnly: true,
		Secure:   true,
		MaxAge:   3600,
	})

}
