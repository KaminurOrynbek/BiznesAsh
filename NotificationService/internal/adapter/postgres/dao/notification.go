package dao

import "time"

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
