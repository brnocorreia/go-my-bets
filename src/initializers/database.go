package initializers

import (
	"fmt"
	"github.com/brnocorreia/go-my-bets/src/configuration/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func ConnectToDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Password, config.User, config.DBName, config.SSLMode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error trying to connect to database")
	}
	logger.Info("Database connected successfully")
}
