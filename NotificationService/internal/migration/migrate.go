package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

func ConnectAndMigrate() *sqlx.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set in environment variables")
	}

	// Connect using sqlx
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set database pool settings
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	fmt.Println("Database connected successfully ✅")

	// Migrations — later we'll run real SQL files separately if needed
	return db
}
