package impl

import (
	"context"
	"fmt"
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

func (u *notificationUsecase) NotifyPostLike(ctx context.Context, n *entity.Notification) error {
	return u.saveTypedNotification(ctx, n, "POST_LIKE")
}

func (u *notificationUsecase) NotifyCommentLike(ctx context.Context, n *entity.Notification) error {
	return u.saveTypedNotification(ctx, n, "COMMENT_LIKE")
}

func (u *notificationUsecase) saveTypedNotification(ctx context.Context, n *entity.Notification, typ string) error {
	// Validate user exists
	exists, err := u.repo.UserExists(ctx, n.UserID)
	if err != nil {
		return fmt.Errorf("failed to verify user: %w", err)
	}
	if !exists {
		return fmt.Errorf("user with ID %s does not exist", n.UserID)
	}

	//Validate post exists
	if n.PostID != nil && *n.PostID != "" {
		exists, err := u.repo.PostExists(ctx, *n.PostID)
		if err != nil {
			return fmt.Errorf("failed to verify post: %w", err)
		}
		if !exists {
			return fmt.Errorf("post with ID %s does not exist", n.PostID)
		}
	}

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

func (u *notificationUsecase) GetWelcomeEmailHTML() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Welcome to BiznesAsh</title>
</head>
<body style="font-family: Arial, sans-serif; background-color: #b3c8e8; padding: 20px;">
  <div style="max-width: 600px; margin: auto; background-color: white; padding: 40px; border-radius: 16px;">
    <h2 style="color: #333333; text-align: left;">Welcome to <span style="color: #003087;">BiznesAsh</span>!</h2>
    <p style="font-size: 16px; color: #333; line-height: 1.5;">Dear Future Entrepreneur,</p>
    <p style="font-size: 16px; color: #333; line-height: 1.5;">
      We are delighted to welcome you to BiznesAsh. Our platform is designed to foster connections, collaboration, and growth among entrepreneurs. We are committed to providing the resources and support you need to thrive in your entrepreneurial journey.
    </p>
    <p style="font-size: 16px; color: #333; line-height: 1.5;">
      We look forward to seeing you leverage the opportunities available within our community.
    </p>
    <p style="font-size: 16px; color: #333; line-height: 1.5;">Sincerely,<br>The BiznesAsh Team</p>
    <div style="display: flex; align-items: center; margin-top: 30px;">
      <img src="https://i.imgur.com/iAPmKNf.jpeg" alt="BiznesAsh Logo" style="width: 80px; height: 80px; margin-right: 20px;">
      <div>
        <p style="font-size: 16px; color: #003087; margin: 0; font-weight: bold;">BiznesAsh</p>
        <p style="font-size: 14px; color: #003087; margin: 5px 0;">biznesash@info.com</p>
        <a href="https://www.biznesash.com" style="font-size: 14px; color: #003087; text-decoration: underline;">www.biznesash.com</a>
      </div>
    </div>
    <div style="text-align: center; margin-top: 20px;">
      <a href="https://www.biznesash.com" style="display: inline-block; background-color: #003087; color: white; padding: 12px 24px; text-decoration: none; border-radius: 24px; font-size: 16px; text-transform: uppercase;">Visit BiznesAsh</a>
    </div>
    <p style="margin-top: 20px; font-size: 12px; color: #666; text-align: center;">If you have any questions, just reply to this email. We're here to help you succeed!</p>
  </div>
</body>
</html>`
}
