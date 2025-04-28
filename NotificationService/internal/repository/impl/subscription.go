package impl

import (
	"context"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	"github.com/jmoiron/sqlx"
)

type subscriptionRepositoryImpl struct {
	db *sqlx.DB
}

func NewSubscriptionRepository(db *sqlx.DB) repo.SubscriptionRepository {
	return &subscriptionRepositoryImpl{db: db}
}

func (r *subscriptionRepositoryImpl) GetSubscriptions(ctx context.Context, userID string) ([]string, error) {
	var subscriptions []string
	query := `SELECT event_type FROM subscriptions WHERE user_id = $1`
	err := r.db.SelectContext(ctx, &subscriptions, query, userID)
	return subscriptions, err
}

func (r *subscriptionRepositoryImpl) AddSubscription(ctx context.Context, userID, eventType string) error {
	query := `INSERT INTO subscriptions (user_id, event_type) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, userID, eventType)
	return err
}

func (r *subscriptionRepositoryImpl) RemoveSubscription(ctx context.Context, userID, eventType string) error {
	query := `DELETE FROM subscriptions WHERE user_id = $1 AND event_type = $2`
	_, err := r.db.ExecContext(ctx, query, userID, eventType)
	return err
}
