package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	
	"google.golang.org/grpc"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/KaminurOrynbek/BiznesAsh/handler"
	contentpb "github.com/KaminurOrynbek/BiznesAsh_lib/proto/auto-proto/content"
	// notificationpb "github.com/KaminurOrynbek/BiznesAsh_lib/proto/auto-proto/notification"
	// userpb "github.com/KaminurOrynbek/BiznesAsh_lib/proto/auto-proto/user"
)


var banMismatchCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "user_ban_mismatch_total",
		Help: "Counts how often a user is marked for ban but their status remains false in DB.",
	},
	[]string{"user_id"},
)


func init() {
	prometheus.MustRegister(banMismatchCounter)
}

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found or error loading .env file")
    }
	// userConn, err := grpc.Dial(os.Getenv("USER_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Failed to connect to UserService: %v", err)
	// }
	
	contentConn, err := grpc.Dial(os.Getenv("CONTENT_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to ContentService: %v", err)
	}
	// notificationConn, err := grpc.Dial(os.Getenv("NOTIFICATION_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Failed to connect to NotificationService: %v", err)
	// }

	// userClient := userpb.NewUserServiceClient(userConn)
	contentClient := contentpb.NewContentServiceClient(contentConn)
	// notificationClient := notificationpb.NewNotificationServiceClient(notificationConn)

	router := gin.Default()

	// handler.RegisterUserRoutes(router, userClient)
	handler.RegisterContentRoutes(router, contentClient)
	// handler.RegisterNotificationRoutes(router, notificationClient)


	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Printf("REST API started at http://localhost:%s", os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
