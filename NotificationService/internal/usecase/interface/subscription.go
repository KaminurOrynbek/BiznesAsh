package _interface

import "context"

type SubscriptionUsecase interface {
	Subscribe(ctx context.Context, userID string, eventTypes []string) error
	Unsubscribe(ctx context.Context, userID string, eventType string) error
	GetSubscriptions(ctx context.Context, userID string) ([]string, error)
}
