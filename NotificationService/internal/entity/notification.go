package entity

import "time"

type Notification struct {
	ID        string
	UserID    string
	Message   string //content
	PostID    *string
	CommentID *string
	Type      string // Type: WELCOME, COMMENT, REPORT, SYSTEM
	CreatedAt time.Time
	IsRead    bool //"read/unread" status.
}
