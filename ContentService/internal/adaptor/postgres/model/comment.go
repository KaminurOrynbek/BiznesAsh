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

func (m *Comment) ToEntity() *entity.Comment {
	return &entity.Comment{
		ID:        m.ID,
		PostID:    m.PostID,
		AuthorID:  m.AuthorID,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
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
