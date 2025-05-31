package subscriber

import (
	"context"
	"encoding/json"
	"log"

	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats/payloads"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/KaminurOrynbek/BiznesAsh_lib/queue"
)

func InitUserSubscribers(q queue.MessageQueue, uc _interface.CombinedUsecase) {
	subscribe := func(subject string, handler func(context.Context, payloads.UserEventPayload)) {
		err := q.Subscribe(subject, func(data []byte) {
			var payload payloads.UserEventPayload
			if err := json.Unmarshal(data, &payload); err != nil {
				log.Printf("❌ Failed to parse payload for %s: %v", subject, err)
				return
			}
			handler(context.Background(), payload)
		})
		if err != nil {
			log.Printf("❌ Failed to subscribe to %s: %v", subject, err)
		}
	}

	// Handle user registration
	subscribe("user.registered", func(ctx context.Context, payload payloads.UserEventPayload) {
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

	// Handle user deletion
	subscribe("user.deleted", func(ctx context.Context, payload payloads.UserEventPayload) {
		_ = uc.SendEmail(ctx, &entity.Email{
			To:      payload.Email,
			Subject: "Account Deletion Confirmation",
			Body:    "Your account has been deleted successfully.",
		})
	})

	// Handle promotions
	subscribe("user.promoted_to_moderator", func(ctx context.Context, payload payloads.UserEventPayload) {
		_ = uc.NotifySystemMessage(ctx, &entity.Notification{
			UserID:  payload.UserID,
			Message: "You were promoted to Moderator.",
		})
	})

	subscribe("user.promoted_to_admin", func(ctx context.Context, payload payloads.UserEventPayload) {
		_ = uc.NotifySystemMessage(ctx, &entity.Notification{
			UserID:  payload.UserID,
			Message: "You were promoted to Admin.",
		})
	})

	// Handle demotions
	subscribe("user.demoted", func(ctx context.Context, payload payloads.UserEventPayload) {
		_ = uc.NotifySystemMessage(ctx, &entity.Notification{
			UserID:  payload.UserID,
			Message: "You have been demoted to User.",
		})
	})

	// Handle bans
	subscribe("user.banned", func(ctx context.Context, payload payloads.UserEventPayload) {
		_ = uc.SendEmail(ctx, &entity.Email{
			To:      payload.Email,
			Subject: "Account Banned",
			Body:    "Your account has been banned for the following reason: " + payload.Reason,
		})
	})
}
