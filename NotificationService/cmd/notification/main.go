package main

import (
	"fmt"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats/subscriber"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	natscfg "github.com/KaminurOrynbek/BiznesAsh/internal/config/nats"
	postgres2 "github.com/KaminurOrynbek/BiznesAsh/internal/config/postgres"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/notification"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres"
	delivery "github.com/KaminurOrynbek/BiznesAsh/internal/delivery/grpc"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/impl"
	usecaseImpl "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/impl"
	"github.com/KaminurOrynbek/BiznesAsh/pkg/queue"
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

	// Load Postgres config
	cfg := postgres2.LoadConfig()
	db := postgres2.ConnectAndMigrate()

	// NATS setup
	natsConfig := natscfg.LoadConfig()
	natsConn := nats.NewConnection(natsConfig)
	defer natsConn.Close()

	notificationDAO := dao.NewNotificationDAO(db)
	subscriptionDAO := dao.NewSubscriptionDAO(db)
	verificationDAO := dao.NewVerificationDAO(db)

	emailSender := postgres.NewEmailSender(cfg)

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

	// Subscribe to NATS events
	subscriber.InitUserSubscribers(natsQueue, combined)

	// Close NATS connection gracefully
	defer func() {
		if err := natsQueue.Close(); err != nil {
			log.Println("Error closing NATS:", err)
		}
	}()

	// Subscribe to "user.registered" subject
	err = natsQueue.Subscribe("user.registered", func(data []byte) {
		fmt.Println("Received user.registered:", string(data))
	})

	if err != nil {
		log.Fatalf("Error subscribing to NATS subject: %v", err)
	}

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register NotificationService server
	pb.RegisterNotificationServiceServer(grpcServer, delivery.NewNotificationDelivery(combined))

	log.Printf("NotificationService gRPC server started on port %s", cfg.GrpcPort)

	// Start serving gRPC requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
