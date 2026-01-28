package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *entity.Comment) error
	Update(ctx context.Context, comment *entity.Comment) error
	Delete(ctx context.Context, commentID string) error
	ListByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)
}
