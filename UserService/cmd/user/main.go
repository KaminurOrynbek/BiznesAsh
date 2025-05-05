package main

import (
	pb "github.com/KaminurOrynbek/BiznesAsh/UserService/auto-proto/user"
	nats "github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/nats"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/postgres/dao"
	natscfg "github.com/KaminurOrynbek/BiznesAsh/UserService/internal/configs/nats"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/configs/posgres"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/delivery/grpc"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/middleware"
	usecase "github.com/KaminurOrynbek/BiznesAsh/UserService/internal/usecase/Impl"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	gogrpc "google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	cfg := posgres.LoadConfig()

	db, err := sqlx.Connect("postgres", cfg.GetDBURL())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := dao.NewUserDAO(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userServer := grpc.NewUserServer(userUsecase)

	grpcServer := gogrpc.NewServer(
		gogrpc.UnaryInterceptor(middleware.AuthInterceptor),
	)

	pb.RegisterUserServiceServer(grpcServer, userServer)

	natsConfig := natscfg.LoadConfig()
	natsConn := nats.NewConnection(natsConfig)
	defer natsConn.Close()

	listener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server listening on port %s", cfg.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
