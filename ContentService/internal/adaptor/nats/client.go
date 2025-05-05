package nats

import (
	natscfg "github.com/KaminurOrynbek/BiznesAsh/internal/config/nats"
	"github.com/nats-io/nats.go"
	"log"
)

// NewConnection создает и возвращает новое подключение к NATS
func NewConnection(cfg *natscfg.Config) *nats.Conn {
	opts := []nats.Option{
		nats.MaxReconnects(cfg.NATSMaxReconnects),
		nats.Timeout(cfg.NATSTimeout),
	}

	conn, err := nats.Connect(cfg.NATSURL, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	log.Printf("Connected to NATS at %s", cfg.NATSURL)
	return conn
}
