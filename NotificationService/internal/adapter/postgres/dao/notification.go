package dao

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/jmoiron/sqlx"
)

type NotificationDAO struct {
	db *sqlx.DB
}

func NewNotificationDAO(db *sqlx.DB) *NotificationDAO {
	return &NotificationDAO{db: db}
}

func (dao *NotificationDAO) Save(ctx context.Context, n *model.Notification) error {
	query := `INSERT INTO notifications (id, user_id, message, post_id, comment_id, type, created_at, is_read)
	          VALUES (:id, :user_id, :message, :post_id, :comment_id, :type, :created_at, :is_read)`
	_, err := dao.db.NamedExecContext(ctx, query, n)
	return err
}
