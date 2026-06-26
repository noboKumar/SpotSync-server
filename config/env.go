package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	dsn       string
	port      string
	jwtSecret string
}

func loadEnv() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		dsn:       os.Getenv("DSN"),
		port:      os.Getenv("PORT"),
		jwtSecret: os.Getenv("JWT_SECRET"),
	}

}
