package main

import (
	"log"
	"net/http"
	"APIGateway/internal/config"
	"ApiGateway/internal/delivery/http"
)

func main() {
	cfg := config.LoadConfig()
	router := http.NewRouter(cfg)
	log.Printf("API Gateway running on %s\n", cfg.ServerAddr)
	log.Fatal(http.ListenAndServe(cfg.ServerAddr, router))
}
