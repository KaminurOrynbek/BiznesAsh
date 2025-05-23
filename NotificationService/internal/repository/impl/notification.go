package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
)

type notificationRepositoryImpl struct {
	dao *dao.NotificationDAO
}

func NewNotificationRepository(dao *dao.NotificationDAO) repo.NotificationRepository {
	return &notificationRepositoryImpl{dao: dao}
}

func (r *notificationRepositoryImpl) SaveNotification(ctx context.Context, notification *entity.Notification) error {
	return r.dao.Save(ctx, model.FromEntityNotification(notification))
}

func (r *notificationRepositoryImpl) UserExists(ctx context.Context, userID string) (bool, error) {
	return r.dao.UserExists(ctx, userID)
}

func (r *notificationRepositoryImpl) PostExists(ctx context.Context, postID string) (bool, error) {
	return r.dao.PostExists(ctx, postID)
}
