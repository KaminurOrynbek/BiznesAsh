package model

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity/enum"
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID            uuid.UUID     `db:"id"`
	Title         string        `db:"title"`
	Content       string        `db:"content"`
	Type          enum.PostType `db:"type"`
	AuthorID      string        `db:"author_id"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
	Published     bool          `db:"published"`
	LikesCount    int32         `db:"likes_count"`
	DislikesCount int32         `db:"dislikes_count"`
	CommentsCount int32         `db:"comments_count"`
	Comments      []*Comment
}

func (Post) TableName() string {
	return "posts"
}

func (p *Post) ToEntity() *entity.Post {
	var entityComments []*entity.Comment
	for _, c := range p.Comments {
		entityComments = append(entityComments, c.ToEntity())
	}

	return &entity.Post{
		ID:            p.ID.String(),
		Title:         p.Title,
		Content:       p.Content,
		Type:          p.Type,
		AuthorID:      p.AuthorID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		Published:     p.Published,
		LikesCount:    p.LikesCount,
		DislikesCount: p.DislikesCount,
		CommentsCount: p.CommentsCount,
		Comments:      entityComments,
	}
}

func FromEntityPost(p *entity.Post) *Post {
	uid, _ := uuid.Parse(p.ID)

	return &Post{
		ID:            uid,
		Title:         p.Title,
		Content:       p.Content,
		Type:          p.Type,
		AuthorID:      p.AuthorID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		Published:     p.Published,
		LikesCount:    p.LikesCount,
		DislikesCount: p.DislikesCount,
		CommentsCount: p.CommentsCount,
	}
}
