package main

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats/subscriber"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/KaminurOrynbek/BiznesAsh_lib/adapter/nats"
	natscfg "github.com/KaminurOrynbek/BiznesAsh_lib/config/nats"
	postgresCfg "github.com/KaminurOrynbek/BiznesAsh_lib/config/postgres"
	"github.com/KaminurOrynbek/BiznesAsh_lib/config/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"

	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/notification"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres"
	delivery "github.com/KaminurOrynbek/BiznesAsh/internal/delivery/grpc"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/impl"
	usecaseImpl "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/impl"

	"github.com/KaminurOrynbek/BiznesAsh_lib/queue"
)

type combinedUsecase struct {
	_interface.NotificationUsecase
	_interface.VerificationUsecase
	_interface.SubscriptionUsecase
	_interface.EmailSender
}

func main() {
	// Load .env if exists
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found, continuing...")
	}

	// Init Postgres
	pgConfig := postgresCfg.LoadPostgresConfig()
	db, err := sqlx.Connect("postgres", pgConfig.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to Postgres!")

	// NATS setup
	natsConfig := natscfg.LoadNatsConfig()
	natsConn := nats.NewConnection(natsConfig)
	defer natsConn.Close()

	notificationDAO := dao.NewNotificationDAO(db)
	subscriptionDAO := dao.NewSubscriptionDAO(db)
	verificationDAO := dao.NewVerificationDAO(db)

	serviceConfig := service.LoadServiceConfig()
	emailSender := postgres.NewEmailSender(serviceConfig)

	notificationRepo := repo.NewNotificationRepository(notificationDAO)
	subscriptionRepo := repo.NewSubscriptionRepository(subscriptionDAO)
	verificationRepo := repo.NewVerificationRepository(verificationDAO)

	combined := &combinedUsecase{
		NotificationUsecase: usecaseImpl.NewNotificationUsecase(notificationRepo, emailSender),
		VerificationUsecase: usecaseImpl.NewVerificationUsecase(verificationRepo, emailSender),
		SubscriptionUsecase: usecaseImpl.NewSubscriptionUsecase(subscriptionRepo),
		EmailSender:         emailSender,
	}

	// NATS Queue initialization
	natsQueue := queue.NewNATSQueue(natsConn)

	// Content Subscriber Setup
	contentSubscriber := subscriber.NewContentSubscriber(natsQueue, combined.NotificationUsecase)

	_ = contentSubscriber.SubscribePostCreated()
	_ = contentSubscriber.SubscribePostUpdated()
	_ = contentSubscriber.SubscribeCommentCreated()
	_ = contentSubscriber.SubscribePostReported()

	// NEW Methods
	_ = contentSubscriber.SubscribePostLiked()
	_ = contentSubscriber.SubscribeCommentLiked()

	// Subscribe to NATS events
	subscriber.InitUserSubscribers(natsQueue, combined)

	// Close NATS connection gracefully
	defer func() {
		if err := natsQueue.Close(); err != nil {
			log.Println("Error closing NATS:", err)
		}
	}()

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		log.Fatal("GRPC_PORT is not set in environment variables")
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register NotificationService server
	pb.RegisterNotificationServiceServer(grpcServer, delivery.NewNotificationDelivery(combined))

	log.Printf("NotificationService gRPC server started on port %s", grpcPort)

	// Start serving gRPC requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
