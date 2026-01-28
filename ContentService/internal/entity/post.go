package entity

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity/enum"
	"time"
)

type Post struct {
	ID            string
	Title         string
	Content       string
	Type          enum.PostType
	AuthorID      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Published     bool
	LikesCount    int32
	DislikesCount int32
	CommentsCount int32
	Comments      []*Comment
}
