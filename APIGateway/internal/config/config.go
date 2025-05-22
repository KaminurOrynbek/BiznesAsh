package config

import "os"

type Config struct {
	ServerAddr             string
	UserServiceURL         string
	ContentServiceURL      string
	NotificationServiceURL string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddr:             getEnv("GATEWAY_ADDR", ":8080"),
		UserServiceURL:         getEnv("USER_SERVICE_URL", "http://localhost:8081"),
		ContentServiceURL:      getEnv("CONTENT_SERVICE_URL", "http://localhost:8082"),
		NotificationServiceURL: getEnv("NOTIFICATION_SERVICE_URL", "http://localhost:8083"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
