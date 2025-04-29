package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type PostRepository interface {
	Create(ctx context.Context, post *entity.Post) error
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*entity.Post, error)
	List(ctx context.Context, offset, limit int) ([]*entity.Post, error)
	Search(ctx context.Context, keyword string, offset, limit int) ([]*entity.Post, error)
}
