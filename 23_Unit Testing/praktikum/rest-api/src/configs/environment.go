package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port             string
	DatabaseUsername string
	DatabasePassword string
	DatabaseHostname string
	DatabasePort     string
	DatabaseName     string
	JwtSecretKey     string
}

func GetServerConfig() ServerConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}

	config := ServerConfig{
		Port:             os.Getenv("PORT"),
		DatabaseUsername: os.Getenv("DB_USERNAME"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
		DatabaseHostname: os.Getenv("DB_HOST"),
		DatabasePort:     os.Getenv("DB_PORT"),
		DatabaseName:     os.Getenv("DB_NAME"),
		JwtSecretKey:     os.Getenv("JWT_SECRET"),
	}
	return config
}
