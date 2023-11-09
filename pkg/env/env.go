package env

import (
	"github.com/joho/godotenv"
	"log"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed loading .env file: %v", err)
	}
}
