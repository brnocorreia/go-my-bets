package main

import (
	"github.com/brnocorreia/go-my-bets/src/initializers"
	"os"
)

func init() {
	initializers.LoadEnvVariables()

	config := &initializers.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	initializers.ConnectToDB(config)

}

func main() {

}
