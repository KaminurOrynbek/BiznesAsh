package entity

import "time"

type Like struct {
	ID        string
	PostID    string
	UserID    string
	IsLike    bool
	CreatedAt time.Time
}
