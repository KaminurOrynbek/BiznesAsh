package impl

import (
	"context"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
)

type subscriptionUsecase struct {
	repo _interface.SubscriptionRepository
}

func NewSubscriptionUsecase(repo _interface.SubscriptionRepository) *subscriptionUsecase {
	return &subscriptionUsecase{repo: repo}
}

func (u *subscriptionUsecase) Subscribe(ctx context.Context, userID string, eventType []string) error {
	for _, eventType := range eventType {
		if err := u.repo.AddSubscription(ctx, userID, eventType); err != nil {
			return err
		}
	}
	return nil
}

func (u *subscriptionUsecase) Unsubscribe(ctx context.Context, userID, eventType string) error {
	return u.repo.RemoveSubscription(ctx, userID, eventType)
}

func (u *subscriptionUsecase) GetSubscriptions(ctx context.Context, userID string) ([]string, error) {
	return u.repo.GetSubscriptions(ctx, userID)
}
