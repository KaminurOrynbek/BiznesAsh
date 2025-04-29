package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type CommentUsecase interface {
	CreateComment(ctx context.Context, comment *entity.Comment) error
	UpdateComment(ctx context.Context, comment *entity.Comment) error
	DeleteComment(ctx context.Context, commentID string) error
	ListCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)
}
