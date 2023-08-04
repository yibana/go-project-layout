package configs

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	// load .env file first
	_ = godotenv.Load()

	PROXY     = os.Getenv("PROXY")
	LOG_DIR   = GetEnvWithDefault("LOG_DIR", "./logs")
	LOG_FILE  = GetEnvWithDefault("LOG_FILE", "default")
	LOG_LEVEL = GetEnvWithDefault("LOG_LEVEL", "debug")

	LOGGER = NewLogger(LOG_DIR, LOG_FILE, LOG_LEVEL)
)

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
