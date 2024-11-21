package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Println("No .env file found. Continuing with existing environment variables.")
		return
	}

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
		// Don't use log.Fatal here as it will stop the entire application
	}
}
