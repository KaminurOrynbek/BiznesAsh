package dao

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity/enum"
	"time"
)

type Post struct {
	ID            string        `db:"id"`
	Title         string        `db:"title"`
	Content       string        `db:"content"`
	Type          enum.PostType `db:"type"`
	AuthorID      string        `db:"author_id"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
	Published     bool          `db:"published"`
	LikesCount    int32         `db:"likes_count"`
	DislikesCount int32         `db:"dislikes_count"`
}

func (Post) TableName() string {
	return "posts"
}
