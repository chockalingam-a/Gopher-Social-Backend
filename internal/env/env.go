package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVar struct {
	ADDR              string
	DB_ADDR           string
	DB_MAX_OPEN_CONNS int
	DB_MAX_IDLE_CONNS int
	DB_MAX_IDLE_TIME  string
}

func LoadEnvVar() EnvVar {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using default or system environment variables")
	}

	return EnvVar{
		ADDR:              GetEnv("ADDR", "3000"),
		DB_ADDR:           GetEnv("DB_ADDR", ""),
		DB_MAX_OPEN_CONNS: GetEnvAsInt("DB_MAX_OPEN_CONNS", 30),
		DB_MAX_IDLE_CONNS: GetEnvAsInt("DB_MAX_IDLE_CONNS", 30),
		DB_MAX_IDLE_TIME:  GetEnv("DB_MAX_IDLE_TIME", "15m"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	valueStr := GetEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Invalid integer for key %s: %s. Using default: %d\n", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}

func GetEnvAsBool(key string, defaultValue bool) bool {
	valueStr := GetEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		log.Printf("Invalid boolean for key %s: %s. Using default: %v\n", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}
