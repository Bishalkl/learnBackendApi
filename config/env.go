package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// type for config
type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWT_SECRET string
	JWTIssure  string
}

// makeing global variable
var Envs = initconfig()

// function initconfig
func initconfig() *Config {

	// directly load in to env
	_ = godotenv.Load()
	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", ":8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "testDb"),
		JWT_SECRET: getEnv("JWT_SECRET", "default"),
		JWTIssure:  getEnv("JWT_ISSUER", "http://localhost:8080"),
	}
}

// func for getEnv
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// // func for getIntEnv
// func getIntEnv(key string, fallback int32) int32 {
// 	if value, ok := os.LookupEnv(key); ok {
// 		// convert string value to int32
// 		parsedValue, err := strconv.Atoi(value)
// 		if err != nil {
// 			return fallback
// 		}
// 		return int32(parsedValue)
// 	}
// 	return fallback
// }
