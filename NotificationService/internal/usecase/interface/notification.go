package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type NotificationUsecase interface {
	SendCommentNotification(ctx context.Context, n *entity.Notification) error
	SendReportNotification(ctx context.Context, n *entity.Notification) error
	NotifyNewPost(ctx context.Context, n *entity.Notification) error
	NotifyPostUpdate(ctx context.Context, n *entity.Notification) error
	NotifySystemMessage(ctx context.Context, n *entity.Notification) error
	GetWelcomeEmailHTML() string
	NotifyPostLike(ctx context.Context, n *entity.Notification) error
	NotifyCommentLike(ctx context.Context, n *entity.Notification) error
}
