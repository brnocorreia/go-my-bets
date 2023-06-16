package initializers

import (
	"github.com/brnocorreia/go-my-bets/src/configuration/logger"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
