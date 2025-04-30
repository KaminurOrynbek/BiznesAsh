package model

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"time"
)

type Comment struct {
	ID        string    `db:"id"`
	PostID    string    `db:"post_id"`
	AuthorID  string    `db:"author_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) ToEntity() *entity.Comment {
	return &entity.Comment{
		ID:        c.ID,
		PostID:    c.PostID,
		AuthorID:  c.AuthorID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func FromEntityComment(e *entity.Comment) *Comment {
	return &Comment{
		ID:        e.ID,
		PostID:    e.PostID,
		AuthorID:  e.AuthorID,
		Content:   e.Content,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
