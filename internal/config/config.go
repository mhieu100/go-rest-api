package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBUrl   string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		AppPort: os.Getenv("PORT"),
		DBUrl:   os.Getenv("DBUrl"),
	}
}
