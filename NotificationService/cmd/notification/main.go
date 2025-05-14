package main

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats"
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
)

type combinedUsecase struct {
	_interface.NotificationUsecase
	_interface.VerificationUsecase
	_interface.SubscriptionUsecase
	_interface.EmailSender
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found, continuing...")
	}

	cfg := postgres2.LoadConfig()
	db := postgres2.ConnectAndMigrate()
	//natsConn := postgres2.ConnectNATS()
	//defer func() {
	//	natsConn.Close()
	//	log.Println("Disconnected from NATS")
	//}()

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

	notificationUC := usecaseImpl.NewNotificationUsecase(notificationRepo, emailSender)
	verificationUC := usecaseImpl.NewVerificationUsecase(verificationRepo, emailSender)
	subscriptionUC := usecaseImpl.NewSubscriptionUsecase(subscriptionRepo)

	combined := &combinedUsecase{
		NotificationUsecase: notificationUC,
		VerificationUsecase: verificationUC,
		SubscriptionUsecase: subscriptionUC,
		EmailSender:         emailSender,
	}

	notificationDelivery := delivery.NewNotificationDelivery(combined)

	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, notificationDelivery)

	log.Printf("NotificationService gRPC server started on port %s ", cfg.GrpcPort)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
