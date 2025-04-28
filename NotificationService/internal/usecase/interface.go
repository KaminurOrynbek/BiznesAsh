// File: internal/usecase/interface.go
package usecase

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

// EmailSender defines behavior for sending emails
type EmailSender interface {
	SendEmail(ctx context.Context, email *entity.Email) error
}

// NotificationRepository defines notification storage behavior
type NotificationRepository interface {
	SaveNotification(ctx context.Context, notification *entity.Notification) error
}

// SubscriptionRepository defines subscription storage behavior
type SubscriptionRepository interface {
	GetSubscriptions(ctx context.Context, userID string) ([]string, error)
	AddSubscription(ctx context.Context, userID, eventType string) error
	RemoveSubscription(ctx context.Context, userID, eventType string) error
}

// VerificationRepository defines verification storage behavior
type VerificationRepository interface {
	SaveVerificationCode(ctx context.Context, verification *entity.Verification) error
	VerifyCode(ctx context.Context, userID, code string) (bool, error)
	GetVerificationCode(ctx context.Context, userID string) (*entity.Verification, error)
	UpdateVerificationStatus(ctx context.Context, userID string) error
	UpdateVerificationCode(ctx context.Context, userID string, newCode string) error
}

// NotificationUsecase defines application logic operations
type NotificationUsecase interface {
	SendWelcomeEmail(ctx context.Context, email *entity.Email) error
	SendVerificationEmail(ctx context.Context, email *entity.Email) error
	ResendVerificationCode(ctx context.Context, userID string) error
	VerifyEmail(ctx context.Context, userID, code string) (bool, error)

	SendCommentNotification(ctx context.Context, notification *entity.Notification) error
	SendReportNotification(ctx context.Context, notification *entity.Notification) error
	NotifyNewPost(ctx context.Context, notification *entity.Notification) error
	NotifyPostUpdate(ctx context.Context, notification *entity.Notification) error
	NotifySystemMessage(ctx context.Context, notification *entity.Notification) error

	Subscribe(ctx context.Context, userID string, eventTypes []string) error
	Unsubscribe(ctx context.Context, userID string, eventType string) error
	GetSubscriptions(ctx context.Context, userID string) ([]string, error)
}
