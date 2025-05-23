package entity

import "time"

type Like struct {
	ID        string
	PostID    string
	UserID    string
	CommentID string
	IsLike    bool
	CreatedAt time.Time
}
