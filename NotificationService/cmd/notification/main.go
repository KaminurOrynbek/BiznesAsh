package main

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"log"
	"net"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/notification"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres"
	"github.com/KaminurOrynbek/BiznesAsh/internal/config"

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

	cfg := config.LoadConfig()
	db := config.ConnectAndMigrate()
	natsConn := config.ConnectNATS()
	defer func() {
		natsConn.Close()
		log.Println("Disconnected from NATS")
	}()

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

	go func() {
		time.Sleep(2 * time.Second)

		conn, err := grpc.Dial("localhost:"+cfg.GrpcPort, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect for test: %v", err)
		}
		defer conn.Close()

		client := pb.NewNotificationServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// 1. SendWelcomeEmail --------------------------------------------------------------------------------
		emailRequest := &pb.EmailRequest{
			Email:   "alimakairat17@gmail.com",
			Subject: "Test: Welcome to BiznesAsh",
			Body: `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Welcome to BiznesAsh</title>
</head>
<body style="font-family: Arial, sans-serif; background-color: #b3c8e8; padding: 20px;">
  <div style="max-width: 600px; margin: auto; background-color: white; padding: 40px; border-radius: 16px;">
    <h2 style="color: #333333; text-align: left;">Welcome to <span style="color: #003087;">BiznesAsh</span>!</h2>
    <p style="font-size: 16px; color: #333; line-height: 1.5;">Dear Future Entrepreneur,</p>
    <p style="font-size: 16px; color: #333; line-height: 1.5;">
      We are delighted to welcome you to BiznesAsh. Our platform is designed to foster connections, collaboration, and growth among entrepreneurs. We are committed to providing the resources and support you need to thrive in your entrepreneurial journey.
    </p>
    <p style="font-size: 16px; color: #333; line-height: 1.5;">
      We look forward to seeing you leverage the opportunities available within our community.
    </p>
    <p style="font-size: 16px; color: #333; line-height: 1.5;">Sincerely,<br>The BiznesAsh Team</p>
    <div style="display: flex; align-items: center; margin-top: 30px;">
      <img src="https://i.imgur.com/iAPmKNf.jpeg" alt="BiznesAsh Logo" style="width: 80px; height: 80px; margin-right: 20px;">
      <div>
        <p style="font-size: 16px; color: #003087; margin: 0; font-weight: bold;">BiznesAsh</p>
        <p style="font-size: 14px; color: #003087; margin: 5px 0;">biznesash@info.com</p>
        <a href="https://www.biznesash.com" style="font-size: 14px; color: #003087; text-decoration: underline;">www.biznesash.com</a>
      </div>
    </div>
    <div style="text-align: center; margin-top: 20px;">
      <a href="https://www.biznesash.com" style="display: inline-block; background-color: #003087; color: white; padding: 12px 24px; text-decoration: none; border-radius: 24px; font-size: 16px; text-transform: uppercase;">Visit BiznesAsh</a>
    </div>
    <p style="margin-top: 20px; font-size: 12px; color: #666; text-align: center;">If you have any questions, just reply to this email. We're here to help you succeed!</p>
  </div>
</body>
</html>`,
		}
		resp, err := client.SendWelcomeEmail(ctx, emailRequest)
		if err != nil {
			log.Fatalf("SendWelcomeEmail failed: %v", err)
		}
		log.Printf("SendWelcomeEmail Response: success=%v, message=%v", resp.Success, resp.Message)

		// 2. SendCommentNotification --------------------------------------------------------------------------------
		commentReq := &pb.CommentNotification{
			Email:       "alimakairat17@gmail.com",
			PostId:      "550e8400-e29b-41d4-a716-446655440000",
			CommentText: "This is a test comment",
		}
		commentResp, err := client.SendCommentNotification(ctx, commentReq)
		if err != nil {
			log.Fatalf("SendCommentNotification failed: %v", err)
		}
		log.Printf("SendCommentNotification Response: success=%v, message=%v", commentResp.Success, commentResp.Message)

		// 3. SendReportNotification --------------------------------------------------------------------------------
		reportReq := &pb.ReportNotification{
			Email:  "alimakairat17@gmail.com",
			PostId: "550e8400-e29b-41d4-a716-446655440000",
			Reason: "Inappropriate Content",
		}
		reportResp, err := client.SendReportNotification(ctx, reportReq)
		if err != nil {
			log.Fatalf("SendReportNotification failed: %v", err)
		}
		log.Printf("SendReportNotification Response: success=%v, message=%v", reportResp.Success, reportResp.Message)

		// 4. NotifyNewPost --------------------------------------------------------------------------------
		newPostReq := &pb.NewPostNotification{
			Email:     "alimakairat17@gmail.com",
			PostTitle: "New Awesome Post",
		}
		newPostResp, err := client.NotifyNewPost(ctx, newPostReq)
		if err != nil {
			log.Fatalf("NotifyNewPost failed: %v", err)
		}
		log.Printf("NotifyNewPost Response: success=%v, message=%v", newPostResp.Success, newPostResp.Message)

		// 5. NotifyPostUpdate --------------------------------------------------------------------------------
		updatePostReq := &pb.PostUpdateNotification{
			Email:         "alimakairat17@gmail.com",
			PostId:        "550e8400-e29b-41d4-a716-446655440000",
			UpdateSummary: "Post title updated",
		}
		updatePostResp, err := client.NotifyPostUpdate(ctx, updatePostReq)
		if err != nil {
			log.Fatalf("NotifyPostUpdate failed: %v", err)
		}
		log.Printf("NotifyPostUpdate Response: success=%v, message=%v", updatePostResp.Success, updatePostResp.Message)

		// 6. SendVerificationEmail --------------------------------------------------------------------------------
		verificationEmailReq := &pb.EmailRequest{
			Email:   "alimakairat17@gmail.com",
			Subject: "Verify Your Email",
			Body:    "",
		}
		verificationResp, err := client.SendVerificationEmail(ctx, verificationEmailReq)
		if err != nil {
			log.Fatalf("SendVerificationEmail failed: %v", err)
		}
		log.Printf("SendVerificationEmail Response: success=%v, message=%v", verificationResp.Success, verificationResp.Message)

		// 7. VerifyEmail --------------------------------------------------------------------------------
		// Перед VerifyEmail --- получаем актуальный код из базы
		var currentCode string
		codeQuery := `SELECT code FROM verifications WHERE user_id = $1 ORDER BY expires_at DESC LIMIT 1`
		err = db.GetContext(ctx, &currentCode, codeQuery, "alimakairat17@gmail.com")
		if err != nil {
			log.Fatalf("Failed to fetch verification code: %v", err)
		}
		// Теперь отправляем правильный код в VerifyEmail
		verifyEmailReq := &pb.VerificationCode{
			Email: "alimakairat17@gmail.com",
			Code:  currentCode, // Must match the code sent in real db if you want successful test
		}
		verifyResp, err := client.VerifyEmail(ctx, verifyEmailReq)
		if err != nil {
			log.Fatalf("VerifyEmail failed: %v", err)
		}
		log.Printf("VerifyEmail Response: success=%v, message=%v", verifyResp.Success, verifyResp.Message)

		// 8. NotifySystemMessage --------------------------------------------------------------------------------
		systemMsgReq := &pb.SystemMessageRequest{
			Email:   "alimakairat17@gmail.com",
			Message: "System maintenance scheduled tonight 11PM 🚀",
		}
		systemMsgResp, err := client.NotifySystemMessage(ctx, systemMsgReq)
		if err != nil {
			log.Fatalf("NotifySystemMessage failed: %v", err)
		}
		log.Printf("NotifySystemMessage Response: success=%v, message=%v", systemMsgResp.Success, systemMsgResp.Message)

		// 9. SubscribeToUpdates --------------------------------------------------------------------------------
		subscribeReq := &pb.UserID{
			UserId: "alimakairat17@gmail.com",
		}
		subResp, err := client.SubscribeToUpdates(ctx, subscribeReq)
		if err != nil {
			log.Fatalf("SubscribeToUpdates failed: %v", err)
		}
		log.Printf("SubscribeToUpdates Response: success=%v, message=%v", subResp.Success, subResp.Message)

		// 10. UnsubscribeFromUpdates --------------------------------------------------------------------------------
		unsubResp, err := client.UnsubscribeFromUpdates(ctx, subscribeReq)
		if err != nil {
			log.Fatalf("UnsubscribeFromUpdates failed: %v", err)
		}
		log.Printf("UnsubscribeFromUpdates Response: success=%v, message=%v", unsubResp.Success, unsubResp.Message)

		// 11. GetSubscriptions --------------------------------------------------------------------------------
		getSubsResp, err := client.GetSubscriptions(ctx, subscribeReq)
		if err != nil {
			log.Fatalf("GetSubscriptions failed: %v", err)
		}
		log.Printf("GetSubscriptions Response: subscriptions=%v", getSubsResp.Subscriptions)

		// 12. ResendVerificationCode --------------------------------------------------------------------------------
		resendCodeResp, err := client.ResendVerificationCode(ctx, subscribeReq)
		if err != nil {
			log.Fatalf("ResendVerificationCode failed: %v", err)
		}
		log.Printf("ResendVerificationCode Response: success=%v, message=%v", resendCodeResp.Success, resendCodeResp.Message)

		//  Fetch the NEW code after Resend
		var resentCode string
		err = db.GetContext(ctx, &resentCode, codeQuery, "alimakairat17@gmail.com")
		if err != nil {
			log.Fatalf("Failed to fetch resent verification code: %v", err)
		}

		// Now verify with the resent code
		resentVerifyReq := &pb.VerificationCode{
			Email: "alimakairat17@gmail.com",
			Code:  resentCode,
		}
		resentVerifyResp, err := client.VerifyEmail(ctx, resentVerifyReq)
		if err != nil {
			log.Fatalf("Verify after resend failed: %v", err)
		}
		log.Printf("Verify after resend Response: success=%v, message=%v", resentVerifyResp.Success, resentVerifyResp.Message)
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
