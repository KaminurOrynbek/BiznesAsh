package model

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"time"
)

type Notification struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Message   string    `db:"message"`
	PostID    *string   `db:"post_id"`    // Pointer to allow NULL
	CommentID *string   `db:"comment_id"` // Pointer to allow NULL
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
	IsRead    bool      `db:"is_read"`
}

func (Notification) TableName() string {
	return "notifications"
}

func (n *Notification) ToEntity() *entity.Notification {
	return &entity.Notification{
		ID:        n.ID,
		UserID:    n.UserID,
		Message:   n.Message,
		PostID:    n.PostID,
		CommentID: n.CommentID,
		Type:      n.Type,
		CreatedAt: n.CreatedAt,
		IsRead:    n.IsRead,
	}
}

func FromEntityNotification(e *entity.Notification) *Notification {
	return &Notification{
		ID:        e.ID,
		UserID:    e.UserID,
		Message:   e.Message,
		PostID:    e.PostID,
		CommentID: e.CommentID,
		Type:      e.Type,
		CreatedAt: e.CreatedAt,
		IsRead:    e.IsRead,
	}
}
