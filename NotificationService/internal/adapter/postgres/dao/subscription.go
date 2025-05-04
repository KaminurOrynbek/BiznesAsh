package dao

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/jmoiron/sqlx"
)

type SubscriptionDAO struct {
	db *sqlx.DB
}

func NewSubscriptionDAO(db *sqlx.DB) *SubscriptionDAO {
	return &SubscriptionDAO{db: db}
}

func (dao *SubscriptionDAO) Add(ctx context.Context, s *model.Subscription) error {
	query := `INSERT INTO subscriptions (id, user_id, event_type) VALUES (:id, :user_id, :event_type)`
	_, err := dao.db.NamedExecContext(ctx, query, s)
	return err
}

func (dao *SubscriptionDAO) Remove(ctx context.Context, userID, eventType string) error {
	query := `DELETE FROM subscriptions WHERE user_id = $1 AND event_type = $2`
	_, err := dao.db.ExecContext(ctx, query, userID, eventType)
	return err
}

func (dao *SubscriptionDAO) List(ctx context.Context, userID string) ([]*model.Subscription, error) {
	var subs []*model.Subscription
	query := `SELECT id, user_id, event_type FROM subscriptions WHERE user_id = $1`
	err := dao.db.SelectContext(ctx, &subs, query, userID)
	return subs, err
}
