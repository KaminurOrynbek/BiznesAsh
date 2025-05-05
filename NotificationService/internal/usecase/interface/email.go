package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type EmailSender interface {
	SendEmail(ctx context.Context, email *entity.Email) error
}
