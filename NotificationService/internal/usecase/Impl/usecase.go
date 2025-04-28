package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"

	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/usecase"
)

type notificationUsecaseImpl struct {
	emailSender            usecase.EmailSender
	notificationRepository usecase.NotificationRepository
	subscriptionRepository usecase.SubscriptionRepository
	verificationRepository usecase.VerificationRepository
}

func NewNotificationUsecase(
	emailSender usecase.EmailSender,
	notificationRepo usecase.NotificationRepository,
	subscriptionRepo usecase.SubscriptionRepository,
	verificationRepo usecase.VerificationRepository,
) usecase.NotificationUsecase {
	return &notificationUsecaseImpl{
		emailSender:            emailSender,
		notificationRepository: notificationRepo,
		subscriptionRepository: subscriptionRepo,
		verificationRepository: verificationRepo,
	}
}

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func (u *notificationUsecaseImpl) SendWelcomeEmail(ctx context.Context, email *entity.Email) error {
	return u.emailSender.SendEmail(ctx, email)
}

func (u *notificationUsecaseImpl) SendVerificationEmail(ctx context.Context, email *entity.Email) error {
	code := generateVerificationCode()

	verification := &entity.Verification{
		UserID:    email.To, // or another unique ID logic depending on your platform
		Email:     email.To,
		Code:      code, // or better generate dynamically
		ExpiresAt: time.Now().Add(10 * time.Minute),
		IsUsed:    false,
	}

	err := u.verificationRepository.SaveVerificationCode(ctx, verification)
	if err != nil {
		return err
	}

	email.Body = "Please verify your account with this code: " + code

	return u.emailSender.SendEmail(ctx, email)
}

func (u *notificationUsecaseImpl) ResendVerificationCode(ctx context.Context, userID string) error {
	v, err := u.verificationRepository.GetVerificationCode(ctx, userID)
	if err != nil {
		return err
	}

	newCode := generateVerificationCode()

	err = u.verificationRepository.UpdateVerificationCode(ctx, userID, newCode)
	if err != nil {
		return err
	}
	email := &entity.Email{
		To:      v.Email,
		Subject: "Resend Verification Code",
		Body:    "Your new verification code is: " + newCode,
	}

	return u.emailSender.SendEmail(ctx, email)
}

func (u *notificationUsecaseImpl) VerifyEmail(ctx context.Context, userID, code string) (bool, error) {
	valid, err := u.verificationRepository.VerifyCode(ctx, userID, code)
	if err != nil || !valid {
		return false, err
	}
	return true, u.verificationRepository.UpdateVerificationStatus(ctx, userID)
}

func (u *notificationUsecaseImpl) SendCommentNotification(ctx context.Context, notification *entity.Notification) error {
	if notification.ID == "" {
		notification.ID = uuid.NewString()
	}
	if notification.CommentID == nil || *notification.CommentID == "" {
		newCommentID := uuid.NewString()
		notification.CommentID = &newCommentID
	}
	if notification.CreatedAt.IsZero() {
		notification.CreatedAt = time.Now()
	}
	notification.Type = "COMMENT"
	notification.IsRead = false
	return u.notificationRepository.SaveNotification(ctx, notification)
}

func (u *notificationUsecaseImpl) SendReportNotification(ctx context.Context, notification *entity.Notification) error {
	if notification.ID == "" {
		notification.ID = uuid.NewString()
	}
	if notification.CreatedAt.IsZero() {
		notification.CreatedAt = time.Now()
	}
	notification.Type = "REPORT"
	notification.IsRead = false
	return u.notificationRepository.SaveNotification(ctx, notification)
}

func (u *notificationUsecaseImpl) NotifyNewPost(ctx context.Context, notification *entity.Notification) error {
	if notification.ID == "" {
		notification.ID = uuid.NewString()
	}
	if notification.CreatedAt.IsZero() {
		notification.CreatedAt = time.Now()
	}
	notification.Type = "NEW_POST"
	notification.IsRead = false
	return u.notificationRepository.SaveNotification(ctx, notification)
}

func (u *notificationUsecaseImpl) NotifyPostUpdate(ctx context.Context, notification *entity.Notification) error {
	if notification.ID == "" {
		notification.ID = uuid.NewString()
	}
	if notification.CreatedAt.IsZero() {
		notification.CreatedAt = time.Now()
	}
	notification.Type = "POST_UPDATE"
	notification.IsRead = false
	return u.notificationRepository.SaveNotification(ctx, notification)
}

func (u *notificationUsecaseImpl) NotifySystemMessage(ctx context.Context, notification *entity.Notification) error {
	if notification.ID == "" {
		notification.ID = uuid.NewString()
	}
	if notification.CreatedAt.IsZero() {
		notification.CreatedAt = time.Now()
	}
	notification.Type = "SYSTEM"
	notification.IsRead = false
	return u.notificationRepository.SaveNotification(ctx, notification)
}

func (u *notificationUsecaseImpl) Subscribe(ctx context.Context, userID string, eventTypes []string) error {
	for _, eventType := range eventTypes {
		if err := u.subscriptionRepository.AddSubscription(ctx, userID, eventType); err != nil {
			return err
		}
	}
	return nil
}

func (u *notificationUsecaseImpl) Unsubscribe(ctx context.Context, userID string, eventType string) error {
	return u.subscriptionRepository.RemoveSubscription(ctx, userID, eventType)
}

func (u *notificationUsecaseImpl) GetSubscriptions(ctx context.Context, userID string) ([]string, error) {
	return u.subscriptionRepository.GetSubscriptions(ctx, userID)
}
