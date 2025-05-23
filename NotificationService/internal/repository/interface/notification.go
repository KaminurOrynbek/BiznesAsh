package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type NotificationRepository interface {
	SaveNotification(ctx context.Context, notification *entity.Notification) error
	UserExists(ctx context.Context, userID string) (bool, error)
	PostExists(ctx context.Context, postID string) (bool, error)
}
