package config

import (
	"log"

	"github.com/joho/godotenv"
)

// GetEnv : getting .env file
func GetEnv() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return env
}
