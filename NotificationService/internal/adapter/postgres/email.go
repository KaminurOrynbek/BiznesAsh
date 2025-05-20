package postgres

import (
	"context"
	"fmt"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/KaminurOrynbek/BiznesAsh_lib/config/service"
	"net/smtp"
)

type smtpEmailSender struct {
	cfg *service.Config
}

func NewEmailSender(cfg *service.Config) _interface.EmailSender {
	return &smtpEmailSender{
		cfg: cfg,
	}
}

func (s *smtpEmailSender) SendEmail(ctx context.Context, email *entity.Email) error {
	auth := smtp.PlainAuth("", s.cfg.SMTPUsername, s.cfg.SMTPPassword, s.cfg.SMTPHost)

	addr := fmt.Sprintf("%s:%s", s.cfg.SMTPHost, s.cfg.SMTPPort)

	msg := []byte(
		"From: " + s.cfg.SMTPUsername + "\n" +
			"To: " + email.To + "\n" +
			"Subject: " + email.Subject + "\n" +
			"MIME-Version: 1.0\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\n\n" +
			email.Body,
	)

	err := smtp.SendMail(addr, auth, s.cfg.SMTPUsername, []string{email.To}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
