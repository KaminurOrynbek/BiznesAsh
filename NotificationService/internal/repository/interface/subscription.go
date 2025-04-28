package _interface

import (
	"context"
)

type SubscriptionRepository interface {
	GetSubscriptions(ctx context.Context, userID string) ([]string, error)
	AddSubscription(ctx context.Context, userID, eventType string) error
	RemoveSubscription(ctx context.Context, userID, eventType string) error
}
