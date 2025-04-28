package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
)

type Config struct {
	GrpcPort          string
	SMTPHost          string
	SMTPPort          string
	SMTPUsername      string
	SMTPPassword      string
	NotificationDBUrl string
	NatsURL           string
}

// LoadConfig reads environment variables and loads them into Config struct
func LoadConfig() *Config {
	return &Config{
		GrpcPort:          getEnv("GRPC_PORT", "50051"),
		SMTPHost:          getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:          getEnv("SMTP_PORT", "587"),
		SMTPUsername:      getEnv("SMTP_USERNAME", ""),
		SMTPPassword:      getEnv("SMTP_PASSWORD", ""),
		NotificationDBUrl: getEnv("NOTIFICATION_DB_URL", "postgres://user:password@localhost:5432/notificationdb?sslmode=disable"),
		NatsURL:           getEnv("NATS_URL", "nats://localhost:4222"),
	}
}

// Helper to get environment variable or fallback default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Connect to PostgreSQL using sqlx
func ConnectAndMigrate() *sqlx.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to database successfully")
	return db
}

// Connect to NATS server

func ConnectNATS() *nats.Conn {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}

	log.Println("Connected to NATS successfully")
	return nc
}
