package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
