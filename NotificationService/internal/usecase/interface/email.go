package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/NotificationService/internal/entity"
)

type EmailSender interface {
	SendEmail(ctx context.Context, email *entity.Email) error
}
