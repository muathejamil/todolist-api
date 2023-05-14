package resources

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnvVariables loads environment variables.
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}
