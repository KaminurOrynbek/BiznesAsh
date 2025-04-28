package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type NotificationRepository interface {
	SaveNotification(ctx context.Context, notification *entity.Notification) error
}
