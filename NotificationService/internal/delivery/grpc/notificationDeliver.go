package grpc

import (
	"context"
	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/notification"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/usecase"
)

type NotificationDelivery struct {
	pb.UnimplementedNotificationServiceServer
	usecase usecase.NotificationUsecase
}

func NewNotificationDelivery(u usecase.NotificationUsecase) *NotificationDelivery {
	return &NotificationDelivery{
		usecase: u,
	}
}

func (d *NotificationDelivery) SendWelcomeEmail(ctx context.Context, req *pb.EmailRequest) (*pb.NotificationResponse, error) {
	email := &entity.Email{
		To:      req.GetEmail(),
		Subject: req.GetSubject(),
		Body:    req.GetBody(),
	}
	err := d.usecase.SendWelcomeEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "Welcome Email Sent"}, nil
}

func (d *NotificationDelivery) SendCommentNotification(ctx context.Context, req *pb.CommentNotification) (*pb.NotificationResponse, error) {
	notification := &entity.Notification{
		UserID:    req.GetEmail(),
		PostID:    ptr(req.GetPostId()),
		CommentID: ptr(req.GetCommentText()),
		Message:   req.GetCommentText(),
		Type:      "COMMENT",
	}
	err := d.usecase.SendCommentNotification(ctx, notification)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "Comment Notification Sent"}, nil
}

func (d *NotificationDelivery) SendReportNotification(ctx context.Context, req *pb.ReportNotification) (*pb.NotificationResponse, error) {
	notification := &entity.Notification{
		UserID:  req.GetEmail(),
		PostID:  ptr(req.GetPostId()),
		Message: req.GetReason(),
		Type:    "REPORT",
	}
	err := d.usecase.SendReportNotification(ctx, notification)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "Report Notification Sent"}, nil
}

func (d *NotificationDelivery) NotifyNewPost(ctx context.Context, req *pb.NewPostNotification) (*pb.NotificationResponse, error) {
	notification := &entity.Notification{
		UserID:  req.GetEmail(),
		Message: req.GetPostTitle(),
		Type:    "NEW_POST",
	}
	err := d.usecase.NotifyNewPost(ctx, notification)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "New Post Notification Sent"}, nil
}

func (d *NotificationDelivery) NotifyPostUpdate(ctx context.Context, req *pb.PostUpdateNotification) (*pb.NotificationResponse, error) {
	notification := &entity.Notification{
		UserID:  req.GetEmail(),
		PostID:  ptr(req.GetPostId()),
		Message: req.GetUpdateSummary(),
		Type:    "POST_UPDATE",
	}
	err := d.usecase.NotifyPostUpdate(ctx, notification)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "Post Update Notification Sent"}, nil
}

func (d *NotificationDelivery) SendVerificationEmail(ctx context.Context, req *pb.EmailRequest) (*pb.NotificationResponse, error) {
	email := &entity.Email{
		To:      req.GetEmail(),
		Subject: req.GetSubject(),
		Body:    req.GetBody(),
	}
	err := d.usecase.SendVerificationEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "Verification Email Sent"}, nil
}

func (d *NotificationDelivery) VerifyEmail(ctx context.Context, req *pb.VerificationCode) (*pb.NotificationResponse, error) {
	success, err := d.usecase.VerifyEmail(ctx, req.GetEmail(), req.GetCode())
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: success, Message: "Verification Result"}, nil
}

func (d *NotificationDelivery) NotifySystemMessage(ctx context.Context, req *pb.SystemMessageRequest) (*pb.NotificationResponse, error) {
	notification := &entity.Notification{
		UserID:  req.GetEmail(),
		Message: req.GetMessage(),
		Type:    "SYSTEM",
	}
	err := d.usecase.NotifySystemMessage(ctx, notification)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "System Message Sent"}, nil
}

func (d *NotificationDelivery) SubscribeToUpdates(ctx context.Context, req *pb.UserID) (*pb.NotificationResponse, error) {
	err := d.usecase.Subscribe(ctx, req.GetUserId(), []string{}) // (you can pass eventTypes if needed)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "Subscribed to updates"}, nil
}

func (d *NotificationDelivery) UnsubscribeFromUpdates(ctx context.Context, req *pb.UserID) (*pb.NotificationResponse, error) {
	err := d.usecase.Unsubscribe(ctx, req.GetUserId(), "") // (you can specify eventType if needed)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "Unsubscribed from updates"}, nil
}

func (d *NotificationDelivery) GetSubscriptions(ctx context.Context, req *pb.UserID) (*pb.SubscriptionsResponse, error) {
	subs, err := d.usecase.GetSubscriptions(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.SubscriptionsResponse{Subscriptions: subs}, nil
}

func (d *NotificationDelivery) ResendVerificationCode(ctx context.Context, req *pb.UserID) (*pb.NotificationResponse, error) {
	err := d.usecase.ResendVerificationCode(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.NotificationResponse{Success: true, Message: "Verification code resent"}, nil
}

// small helper function to create *string
func ptr(s string) *string {
	return &s
}
