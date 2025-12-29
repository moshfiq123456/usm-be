package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         int
	DBDialect    string
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	LogFilePath  string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using environment variables")
	}

	port, _ := strconv.Atoi(getEnv("PORT", "8080"))
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5433"))

	return &Config{
		Port:        port,
		DBDialect:   getEnv("DB_DIALECT", "postgres"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      dbPort,
		DBUser:      getEnv("DB_USER", "go_admin"),
		DBPassword:  getEnv("DB_PASSWORD", "GoAdmin@1234567"),
		DBName:      getEnv("DB_NAME", "go_training"),
		DBSSLMode:   getEnv("DB_SSLMODE", "disable"),
		LogFilePath: getEnv("LOG_FILE_PATH", "./logs/app.log"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}