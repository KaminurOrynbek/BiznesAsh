package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	"github.com/jmoiron/sqlx"
)

type notificationRepositoryImpl struct {
	db *sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) repo.NotificationRepository {
	return &notificationRepositoryImpl{db: db}
}

func (r *notificationRepositoryImpl) SaveNotification(ctx context.Context, notification *entity.Notification) error {
	query := `
		INSERT INTO notifications (id, user_id, message, post_id, comment_id, type, created_at, is_read)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		notification.ID,
		notification.UserID,
		notification.Message,
		notification.PostID,
		notification.CommentID,
		notification.Type,
		notification.CreatedAt,
		notification.IsRead,
	)
	return err
}
