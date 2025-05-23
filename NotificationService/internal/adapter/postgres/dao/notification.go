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

func (dao *NotificationDAO) UserExists(ctx context.Context, userID string) (bool, error) {
	const query = `SELECT 1 FROM users WHERE id = $1 LIMIT 1`
	var exists int
	err := dao.db.GetContext(ctx, &exists, query, userID)
	if err != nil {
		// Not found is not an error, just return false
		return false, nil
	}
	return true, nil
}

func (dao *NotificationDAO) PostExists(ctx context.Context, postID string) (bool, error) {
	const query = `SELECT 1 FROM posts WHERE id = $1 LIMIT 1`
	var exists int
	err := dao.db.GetContext(ctx, &exists, query, postID)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (dao *NotificationDAO) DB() *sqlx.DB {
	return dao.db
}
