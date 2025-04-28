package main

import (
	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/user"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/internal/configs/posgres"
	"github.com/KaminurOrynbek/BiznesAsh/internal/delivery/grpc"
	"github.com/KaminurOrynbek/BiznesAsh/internal/middleware"
	"github.com/KaminurOrynbek/BiznesAsh/internal/usecase/Impl"
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

	listener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server listening on port %s", cfg.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
