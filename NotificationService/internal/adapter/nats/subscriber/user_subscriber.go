package subscriber

import (
	"context"
	"encoding/json"
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

func InitUserSubscribers(q queue.MessageQueue, uc usecase.NotificationUsecase) {
	subscribe(q, "user.registered", func(payload UserEventPayload) {
		_ = uc.SendVerificationEmail(context.Background(), payload.UserID)
		_ = uc.SendWelcomeEmail(context.Background(), payload.UserID)
	})

	subscribe(q, "user.deleted", func(payload UserEventPayload) {
		_ = uc.SendDeactivationEmail(context.Background(), payload.UserID)
	})

	subscribe(q, "user.promoted_to_moderator", func(payload UserEventPayload) {
		_ = uc.NotifySystemMessage(context.Background(), payload.UserID, "You are now a moderator")
	})

	subscribe(q, "user.promoted_to_admin", func(payload UserEventPayload) {
		_ = uc.NotifySystemMessage(context.Background(), payload.UserID, "You are now an admin")
	})

	subscribe(q, "user.demoted", func(payload UserEventPayload) {
		_ = uc.NotifySystemMessage(context.Background(), payload.UserID, "You have been demoted to user")
	})

	subscribe(q, "user.banned", func(payload UserEventPayload) {
		_ = uc.SendBanNotification(context.Background(), payload.UserID, payload.Reason)
	})
}

func subscribe(q queue.MessageQueue, subject string, handler func(UserEventPayload)) {
	err := q.Subscribe(subject, func(data []byte) {
		var payload UserEventPayload
		if err := json.Unmarshal(data, &payload); err != nil {
			log.Printf("Failed to parse %s: %v", subject, err)
			return
		}
		handler(payload)
	})
	if err != nil {
		log.Printf("Failed to subscribe to %s: %v", subject, err)
	}
}
