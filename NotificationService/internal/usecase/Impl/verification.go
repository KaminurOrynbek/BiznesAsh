package impl

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	usecase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
)

type verificationUsecaseImpl struct {
	repo        repo.VerificationRepository
	emailSender usecase.EmailSender
}

func NewVerificationUsecase(
	repo repo.VerificationRepository,
	sender usecase.EmailSender,
) usecase.VerificationUsecase {
	return &verificationUsecaseImpl{
		repo:        repo,
		emailSender: sender,
	}
}

func (u *verificationUsecaseImpl) SendVerificationEmail(ctx context.Context, email *entity.Email) error {
	code := generateVerificationCode()
	verification := &entity.Verification{
		UserID:    email.To,
		Email:     email.To,
		Code:      code,
		ExpiresAt: time.Now().Add(10 * time.Minute),
		IsUsed:    false,
	}

	if err := u.repo.SaveVerificationCode(ctx, verification); err != nil {
		return err
	}

	email.Body = "Please verify your account with this code: " + code
	return u.emailSender.SendEmail(ctx, email)
}

func (u *verificationUsecaseImpl) ResendVerificationCode(ctx context.Context, userID string) error {
	v, err := u.repo.GetVerificationCode(ctx, userID)
	if err != nil {
		return err
	}

	newCode := generateVerificationCode()
	if err := u.repo.UpdateVerificationCode(ctx, userID, newCode); err != nil {
		return err
	}

	email := &entity.Email{
		To:      v.Email,
		Subject: "Resend Verification Code",
		Body:    "Your new verification code is: " + newCode,
	}
	return u.emailSender.SendEmail(ctx, email)
}

func (u *verificationUsecaseImpl) VerifyEmail(ctx context.Context, userID, code string) (bool, error) {
	valid, err := u.repo.VerifyCode(ctx, userID, code)
	if err != nil || !valid {
		return false, err
	}
	return true, u.repo.UpdateVerificationStatus(ctx, userID)
}

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
