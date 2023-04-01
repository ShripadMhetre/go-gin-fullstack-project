package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	fmt.Println("Loading environment variables...")
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// Abstraction for fetching Environment variables
func GetEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal("Environment variable not found for: ", key)
	}

	return value
}
