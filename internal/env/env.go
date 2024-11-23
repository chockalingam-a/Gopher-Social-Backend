package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVar struct {
	PORT string
}

func LoadEnvVar() EnvVar {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using default or system environment variables")
	}

	return EnvVar{
		PORT: GetEnv("PORT", "3000"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
