package posgres

import (
	"fmt"
	"os"
)

// Config holds the configuration values for the application
type Config struct {
	JWTSecret  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	GRPCPort   string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		JWTSecret:  os.Getenv("JWT_SECRET"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		GRPCPort:   os.Getenv("GRPC_PORT"),
	}
}

// GetDBURL combines the DB configuration values into a single connection string
func (c *Config) GetDBURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
