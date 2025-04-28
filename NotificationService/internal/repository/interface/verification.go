package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type VerificationRepository interface {
	SaveVerificationCode(ctx context.Context, verification *entity.Verification) error
	VerifyCode(ctx context.Context, userID, code string) (bool, error)
	GetVerificationCode(ctx context.Context, userID string) (*entity.Verification, error)
	UpdateVerificationStatus(ctx context.Context, userID string) error
	UpdateVerificationCode(ctx context.Context, userID string, newCode string) error
}
