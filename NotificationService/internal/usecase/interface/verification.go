package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type VerificationUsecase interface {
	SendVerificationEmail(ctx context.Context, email *entity.Email) error
	ResendVerificationCode(ctx context.Context, userID string) error
	VerifyEmail(ctx context.Context, userID, code string) (bool, error)
}
