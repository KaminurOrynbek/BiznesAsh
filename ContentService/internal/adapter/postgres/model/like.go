package model

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/google/uuid"
	"time"
)

type Like struct {
	ID        uuid.UUID `db:"id"`
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
		ID:        c.ID.String(),
		PostID:    c.PostID,
		UserID:    c.UserID,
		IsLike:    c.IsLike,
		CreatedAt: c.CreatedAt,
	}
}

func FromEntityLike(e *entity.Like) *Like {
	uid, _ := uuid.Parse(e.ID)
	return &Like{
		ID:        uid,
		PostID:    e.PostID,
		UserID:    e.UserID,
		IsLike:    e.IsLike,
		CreatedAt: e.CreatedAt,
	}
}
