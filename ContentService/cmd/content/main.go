package main

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats"
	natscfg "github.com/KaminurOrynbek/BiznesAsh/internal/config/nats"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	pgcfg "github.com/KaminurOrynbek/BiznesAsh/internal/config/postgres"

	repoimpl "github.com/KaminurOrynbek/BiznesAsh/internal/repository/Impl"
	usecaseimpl "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/impl"

	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/content"
	handler "github.com/KaminurOrynbek/BiznesAsh/internal/delivery/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load")
	} else {
		log.Println("env file loaded successfully")
	}
}

func main() {

	// 1. Load Config
	pgConfig := pgcfg.LoadPostgresConfig()
	//redisConfig := rediscfg.LoadRedisConfig()

	// 2. Init DB
	db := postgres.NewPostgres(pgConfig.DSN())

	//// 3. Init Redis
	//redisClient := redis.NewRedisClient(redisConfig.Addr, redisConfig.Password, redisConfig.DB)
	//if err := redisClient.Ping(context.Background()); err != nil {
	//	log.Fatalf("Redis connection failed: %v", err)
	//}

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

	natsConfig := natscfg.LoadConfig()
	natsConn := nats.NewConnection(natsConfig)
	defer natsConn.Close()

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

}
