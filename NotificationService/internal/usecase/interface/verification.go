package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type VerificationUsecase interface {
	SendVerificationEmail(ctx context.Context, email *entity.Email) error
}
