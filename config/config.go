package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN   string
	AppPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("[CONFIG] No .env file found, falling back to system environment variables")
	}

	return &Config{
		DBDSN:   getDSN(),
		AppPort: getEnv("APP_PORT", "8080"),
	}
}

func getDSN() string {
	return "host=" + getEnv("DB_HOST", "localhost") +
		" user=" + getEnv("DB_USER", "wallet_user") +
		" password=" + getEnv("DB_PASSWORD", "wallet_password") +
		" dbname=" + getEnv("DB_NAME", "digital_wallet") +
		" port=" + getEnv("DB_PORT", "5432") +
		" sslmode=" + getEnv("DB_SSLMODE", "disable")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}