package main

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/pkg/queue"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pgcfg "github.com/KaminurOrynbek/BiznesAsh/internal/config/postgres"
	rediscfg "github.com/KaminurOrynbek/BiznesAsh/internal/config/redis"

	"github.com/KaminurOrynbek/BiznesAsh/internal/adaptor/postgres"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adaptor/redis"

	"github.com/KaminurOrynbek/BiznesAsh/internal/adaptor/postgres/dao"

	repoimpl "github.com/KaminurOrynbek/BiznesAsh/internal/repository/Impl"
	usecaseimpl "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/impl"

	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/content"
	handler "github.com/KaminurOrynbek/BiznesAsh/internal/delivery/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found or failed to load")
	} else {
		log.Println("✅ .env file loaded successfully")
	}
}

func main() {

	// 1. Load Config
	pgConfig := pgcfg.LoadPostgresConfig()
	redisConfig := rediscfg.LoadRedisConfig()

	// 2. Init DB
	db := postgres.NewPostgres(pgConfig.DSN())

	// 3. Init Redis
	redisClient := redis.NewRedisClient(redisConfig.Addr, redisConfig.Password, redisConfig.DB)
	if err := redisClient.Ping(context.Background()); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	// 4. Init DAOs
	postDAO := dao.NewPostDAO(db)
	commentDAO := dao.NewCommentDAO(db)
	likeDAO := dao.NewLikeDAO(db)

	// 5. Init Repositories
	postRepo := repoimpl.NewPostRepository(postDAO)
	commentRepo := repoimpl.NewCommentRepository(commentDAO)
	likeRepo := repoimpl.NewLikeRepository(likeDAO)

	// 6. Init Usecases
	postUsecase := usecaseimpl.NewPostUsecase(postRepo, commentRepo)
	commentUsecase := usecaseimpl.NewCommentUsecase(commentRepo)
	likeUsecase := usecaseimpl.NewLikeUsecase(likeRepo)

	// 7. Init gRPC handler
	contentHandler := handler.NewContentHandler(postUsecase, commentUsecase, likeUsecase)

	// 8. Start gRPC server
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50055" // fallback if not set
	}
	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterContentServiceServer(s, contentHandler)

	log.Println("gRPC server is running on :50055")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	ctx := context.Background()

	natsClient, err := queue.NewClient(ctx, []string{"queue://localhost:4222"}, os.Getenv("NKEY_SEED"), false)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer natsClient.Close()
}
