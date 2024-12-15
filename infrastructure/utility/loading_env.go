package utility

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadingEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
    	log.Println("Error loading .env file")
  	}
}