package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	ServerPort string
	ServerMode string
}

var config Configuration

func init() {
	godotenv.Load()

	config = Configuration{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerMode: getEnv("SERVER_MODE", "debug"),
	}
}

func GetConfig() Configuration {
	return config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
