package redis

import (
	"log"
	"os"
)

// Config holds the configuration values for Redis
type Config struct {
	RedisURL string
}

// LoadConfig loads the Redis configuration from environment variables
func LoadConfig() *Config {
	redisURL := getEnv("REDIS_URL", "localhost:6379")

	return &Config{
		RedisURL: redisURL,
	}
}

// GetRedisURL returns the Redis URL
func (c *Config) GetRedisURL() string {
	return c.RedisURL
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not set, using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
