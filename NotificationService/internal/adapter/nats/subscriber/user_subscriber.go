package subscriber

import (
	"context"
	"encoding/json"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	usecase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"

	"github.com/KaminurOrynbek/BiznesAsh/pkg/queue"
	"log"
)

type UserEventPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email,omitempty"`
	Role   string `json:"role,omitempty"`
	Reason string `json:"reason,omitempty"`
}

func InitUserSubscribers(q queue.MessageQueue, uc usecase.CombinedUsecase) {
	subscribe := func(subject string, handler func(context.Context, UserEventPayload)) {
		err := q.Subscribe(subject, func(data []byte) {
			var payload UserEventPayload
			if err := json.Unmarshal(data, &payload); err != nil {
				log.Printf("Failed to parse payload for %s: %v", subject, err)
				return
			}
			handler(context.Background(), payload)
		})
		if err != nil {
			log.Printf("Failed to subscribe to %s: %v", subject, err)
		}
	}

	subscribe("user.registered", func(ctx context.Context, payload UserEventPayload) {
		_ = uc.SendVerificationEmail(ctx, &entity.Email{
			To:      payload.Email,
			Subject: "Please verify your email",
			Body:    "Welcome! Your verification code will arrive shortly.",
		})
		_ = uc.SendEmail(ctx, &entity.Email{
			To:      payload.Email,
			Subject: "Welcome to BiznesAsh!",
			Body:    uc.GetWelcomeEmailHTML(),
		})
	})

	subscribe("user.deleted", func(ctx context.Context, payload UserEventPayload) {
		_ = uc.SendEmail(ctx, &entity.Email{
			To:      payload.Email,
			Subject: "Account Deletion Confirmation",
			Body:    "Your account has been deleted successfully.",
		})
	})

	subscribe("user.promoted_to_moderator", func(ctx context.Context, payload UserEventPayload) {
		_ = uc.NotifySystemMessage(ctx, &entity.Notification{
			UserID:  payload.UserID,
			Message: "You were promoted to Moderator.",
		})
	})
	subscribe("user.promoted_to_admin", func(ctx context.Context, payload UserEventPayload) {
		_ = uc.NotifySystemMessage(ctx, &entity.Notification{
			UserID:  payload.UserID,
			Message: "You were promoted to Admin.",
		})
	})

	subscribe("user.demoted", func(ctx context.Context, payload UserEventPayload) {
		_ = uc.NotifySystemMessage(ctx, &entity.Notification{
			UserID:  payload.UserID,
			Message: "You have been demoted to User.",
		})
	})

	subscribe("user.banned", func(ctx context.Context, payload UserEventPayload) {
		_ = uc.SendEmail(ctx, &entity.Email{
			To:      payload.Email,
			Subject: "Account Banned",
			Body:    "Your account has been banned for the following reason: " + payload.Reason,
		})
	})
}
