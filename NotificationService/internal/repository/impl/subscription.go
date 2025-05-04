package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
)

type subscriptionRepositoryImpl struct {
	dao *dao.SubscriptionDAO
}

func NewSubscriptionRepository(dao *dao.SubscriptionDAO) repo.SubscriptionRepository {
	return &subscriptionRepositoryImpl{dao: dao}
}

func (r *subscriptionRepositoryImpl) GetSubscriptions(ctx context.Context, userID string) ([]string, error) {
	subs, err := r.dao.List(ctx, userID)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, s := range subs {
		result = append(result, s.EventType) // Assuming EventTypes is a comma-separated string
	}
	return result, nil
}

func (r *subscriptionRepositoryImpl) AddSubscription(ctx context.Context, userID, eventType string) error {
	return r.dao.Add(ctx, &model.Subscription{
		UserID:    userID,
		EventType: eventType,
	})
}

func (r *subscriptionRepositoryImpl) RemoveSubscription(ctx context.Context, userID, eventType string) error {
	return r.dao.Remove(ctx, userID, eventType)
}
