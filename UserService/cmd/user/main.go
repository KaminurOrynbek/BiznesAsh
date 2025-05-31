package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	gogrpc "google.golang.org/grpc"

	pb "github.com/KaminurOrynbek/BiznesAsh/UserService/auto-proto/user"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/nats/publisher"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/delivery/grpc"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/middleware"
	usecase "github.com/KaminurOrynbek/BiznesAsh/UserService/internal/usecase/Impl"
	"github.com/KaminurOrynbek/BiznesAsh_lib/adapter/nats"
	natscfg "github.com/KaminurOrynbek/BiznesAsh_lib/config/nats"
	postgresCfg "github.com/KaminurOrynbek/BiznesAsh_lib/config/postgres"
	"github.com/KaminurOrynbek/BiznesAsh_lib/queue"
)

func main() {
	// Try to load .env file, but don't fail if it doesn't exist
	if err := godotenv.Load(); err != nil {
		// Only log if it's not a "file not found" error
		if !os.IsNotExist(err) {
			log.Printf("Warning: Error loading .env file: %v", err)
		}
	}

	// Init Postgres
	pgConfig := postgresCfg.LoadPostgresConfig()
	db, err := sqlx.Connect("postgres", pgConfig.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to Postgres!")

	// Init user repo
	userRepo := dao.NewUserDAO(db)

	// NATS setup
	natsConfig := natscfg.LoadNatsConfig()
	natsConn := nats.NewConnection(natsConfig)
	defer natsConn.Close()

	// Wrap NATS into queue-compatible interface
	msgQueue := queue.NewNATSQueue(natsConn)

	// Create publisher and usecase
	userPublisher := publisher.NewUserPublisher(msgQueue)
	userUsecase := usecase.NewUserUsecase(userRepo, userPublisher)

	// Get GRPC_PORT from environment
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		log.Fatal("GRPC_PORT is not set in environment variables")
	}

	// Create gRPC server
	userServer := grpc.NewUserServer(userUsecase)
	grpcServer := gogrpc.NewServer(
		gogrpc.UnaryInterceptor(middleware.AuthInterceptor),
	)

	// Register gRPC service
	pb.RegisterUserServiceServer(grpcServer, userServer)

	reflection.Register(grpcServer)

	// Start listener
	listener, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server listening on port %s", grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
