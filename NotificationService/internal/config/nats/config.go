package nats

import (
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
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

func ConnectNATS() *nats.Conn {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "queue://localhost:4222"
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}

	log.Println("Connected to NATS successfully")
	return nc
}
