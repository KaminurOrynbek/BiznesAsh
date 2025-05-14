package main

import (
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	gogrpc "google.golang.org/grpc"

	pb "github.com/KaminurOrynbek/BiznesAsh/UserService/auto-proto/user"
	nats "github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/nats"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/nats/publisher"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/pkg/queue"
	natscfg "github.com/KaminurOrynbek/BiznesAsh/UserService/internal/configs/nats"
	posgres "github.com/KaminurOrynbek/BiznesAsh/UserService/internal/configs/posgres"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/delivery/grpc"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/middleware"
	usecase "github.com/KaminurOrynbek/BiznesAsh/UserService/internal/usecase/Impl"
)

func main() {
	// Load .env if exists
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Load Postgres config
	cfg := posgres.LoadConfig()

	// Connect to Postgres
	db, err := sqlx.Connect("postgres", cfg.GetDBURL())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Init user repo
	userRepo := dao.NewUserDAO(db)

	// NATS setup
	natsConfig := natscfg.LoadConfig()
	rawConn := nats.NewConnection(natsConfig)
	defer rawConn.Close()

	// Wrap NATS into queue-compatible interface
	msgQueue := queue.NewNATSQueue(rawConn)

	// Create publisher and usecase
	userPublisher := publisher.NewUserPublisher(msgQueue)
	userUsecase := usecase.NewUserUsecase(userRepo, userPublisher)

	// Create gRPC server
	userServer := grpc.NewUserServer(userUsecase)
	grpcServer := gogrpc.NewServer(
		gogrpc.UnaryInterceptor(middleware.AuthInterceptor),
	)

	// Register gRPC service
	pb.RegisterUserServiceServer(grpcServer, userServer)

	// Start listener
	listener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server listening on port %s", cfg.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
