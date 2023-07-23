package utils

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	defaultPort = "8080"
)

func LoadConfig() {
	_ = godotenv.Load()
}

func GetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		return defaultPort
	}
	return port
}
