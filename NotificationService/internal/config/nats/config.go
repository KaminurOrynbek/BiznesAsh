package nats

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	NATSURL           string
	NATSMaxReconnects int
	NATSTimeout       time.Duration
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maxReconnects, _ := strconv.Atoi(os.Getenv("NATS_MAX_RECONNECTS"))
	timeout, _ := time.ParseDuration(os.Getenv("NATS_TIMEOUT"))

	return &Config{
		NATSURL:           os.Getenv("NATS_URL"),
		NATSMaxReconnects: maxReconnects,
		NATSTimeout:       timeout,
	}
}
