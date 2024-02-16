package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(envfile string) {
	err := godotenv.Load(envfile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
