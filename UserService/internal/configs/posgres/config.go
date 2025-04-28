package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	GRPCPort   string
	JWTSecret  string
	RedisURL   string
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func LoadConfig() *Config {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "0000")
	dbName := getEnv("DB_NAME", "user_service")
	grpcPort := getEnv("GRPC_PORT", "50051")
	jwtSecret := getEnv("JWT_SECRET", "your_jwt_secret_key_very_long_and_secure_string")
	redisURL := getEnv("REDIS_URL", "localhost:6379")

	return &Config{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		GRPCPort:   grpcPort,
		JWTSecret:  jwtSecret,
		RedisURL:   redisURL,
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not set, using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
