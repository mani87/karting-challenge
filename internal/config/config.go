package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ApiKey string
var CreateOrderApiKey string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	ApiKey = os.Getenv("API_KEY")
	CreateOrderApiKey = os.Getenv("CREATE_ORDER_API_KEY")
	if ApiKey == "" {
		log.Fatal("API_KEY not set")
	}
	if CreateOrderApiKey == "" {
		log.Fatal("CREATE_ORDER_API_KEY not set")
	}
}
