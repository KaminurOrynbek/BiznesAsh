package model

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/google/uuid"
	"time"
)

type Like struct {
	ID        uuid.UUID  `db:"id"`
	PostID    *uuid.UUID `db:"post_id"`
	UserID    string     `db:"user_id"`
	CommentID *uuid.UUID `db:"comment_id"`
	IsLike    bool       `db:"is_like"`
	CreatedAt time.Time  `db:"created_at"`
}

func (Like) TableName() string {
	return "likes"
}

func (c *Like) ToEntity() *entity.Like {
	var postIDStr, commentIDStr string

	if c.PostID != nil {
		postIDStr = c.PostID.String()
	}
	if c.CommentID != nil {
		commentIDStr = c.CommentID.String()
	}

	return &entity.Like{
		ID:        c.ID.String(),
		PostID:    postIDStr,
		UserID:    c.UserID,
		CommentID: commentIDStr,
		IsLike:    c.IsLike,
		CreatedAt: c.CreatedAt,
	}
}

func FromEntityLike(e *entity.Like) *Like {
	id, _ := uuid.Parse(e.ID)

	var postUUID, commentUUID *uuid.UUID

	if e.PostID != "" {
		parsed := uuid.MustParse(e.PostID)
		postUUID = &parsed
	}
	if e.CommentID != "" {
		parsed := uuid.MustParse(e.CommentID)
		commentUUID = &parsed
	}

	return &Like{
		ID:        id,
		PostID:    postUUID,
		UserID:    e.UserID,
		CommentID: commentUUID,
		IsLike:    e.IsLike,
		CreatedAt: e.CreatedAt,
	}
}
