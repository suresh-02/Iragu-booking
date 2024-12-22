package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suresh-02/Iragu-booking/database"
	"github.com/suresh-02/Iragu-booking/models"
)

type contextKey string

const userContextKey contextKey = "user"

// Validate is middleware to validate JWT tokens and retrieve user information.
func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the auth token from cookies
		tokenCookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
			return
		}

		// Extract the token string
		tokenString := tokenCookie.Value

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
			return
		}

		// Check if the token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			http.Error(w, "Session expired. Please login again.", http.StatusUnauthorized)
			return
		}

		// Retrieve user from the database
		var user models.UserCreds
		database.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
			return
		}

		// Store user information in the request context
		ctx := context.WithValue(r.Context(), userContextKey, user)

		// Call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
