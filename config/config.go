package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN   string
	AppPort string
	MaxRetries int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("[CONFIG] No .env file found, falling back to system environment variables")
	}

	return &Config{
		DBDSN:   getDSN(),
		AppPort: getEnv("APP_PORT", "8080"),
		MaxRetries: getEnvAsInt("MAX_RETRIES", 3),
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

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return fallback
}