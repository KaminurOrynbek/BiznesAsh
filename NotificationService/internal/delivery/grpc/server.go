package grpc

import (
	//"net"
	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/notification"
	"google.golang.org/grpc"
)

func NewServer(notificationService pb.NotificationServiceServer) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterNotificationServiceServer(server, notificationService)
	return server
}
