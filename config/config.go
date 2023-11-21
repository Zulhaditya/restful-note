package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	// return the value of variable
	return os.Getenv(key)
}
