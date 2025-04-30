package model

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"time"
)

type Like struct {
	ID        string    `db:"id"`
	PostID    string    `db:"post_id"`
	UserID    string    `db:"user_id"`
	IsLike    bool      `db:"is_like"`
	CreatedAt time.Time `db:"created_at"`
}

func (Like) TableName() string {
	return "likes"
}

func (c *Like) ToEntity() *entity.Like {
	return &entity.Like{
		ID:        c.ID,
		PostID:    c.PostID,
		UserID:    c.UserID,
		IsLike:    c.IsLike,
		CreatedAt: c.CreatedAt,
	}
}

func FromEntityLike(e *entity.Like) *Like {
	return &Like{
		ID:        e.ID,
		PostID:    e.PostID,
		UserID:    e.UserID,
		IsLike:    e.IsLike,
		CreatedAt: e.CreatedAt,
	}
}
