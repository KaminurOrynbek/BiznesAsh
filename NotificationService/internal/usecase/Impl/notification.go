package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	usecase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/google/uuid"
	"time"
)

type notificationUsecase struct {
	repo _interface.NotificationRepository
}

func NewNotificationUsecase(repo _interface.NotificationRepository, sender usecase.EmailSender) *notificationUsecase {
	return &notificationUsecase{repo: repo}
}

func (u *notificationUsecase) SendCommentNotification(ctx context.Context, n *entity.Notification) error {
	return u.saveTypedNotification(ctx, n, "COMMENT")
}

func (u *notificationUsecase) SendReportNotification(ctx context.Context, n *entity.Notification) error {
	return u.saveTypedNotification(ctx, n, "REPORT")
}

func (u *notificationUsecase) NotifyNewPost(ctx context.Context, n *entity.Notification) error {
	return u.saveTypedNotification(ctx, n, "NEW_POST")
}

func (u *notificationUsecase) NotifyPostUpdate(ctx context.Context, n *entity.Notification) error {
	return u.saveTypedNotification(ctx, n, "POST_UPDATE")
}

func (u *notificationUsecase) NotifySystemMessage(ctx context.Context, n *entity.Notification) error {
	return u.saveTypedNotification(ctx, n, "SYSTEM")
}

func (u *notificationUsecase) saveTypedNotification(ctx context.Context, n *entity.Notification, typ string) error {
	if n.ID == "" {
		n.ID = uuid.NewString()
	}
	if n.CreatedAt.IsZero() {
		n.CreatedAt = time.Now()
	}
	n.Type = typ
	n.IsRead = false
	return u.repo.SaveNotification(ctx, n)
}
