package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	ucase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"

	"time"
)

type commentUsecaseImpl struct {
	repo _interface.CommentRepository
}

func NewCommentUsecase(repo _interface.CommentRepository) ucase.CommentUsecase {
	return &commentUsecaseImpl{repo: repo}
}

func (u *commentUsecaseImpl) CreateComment(ctx context.Context, comment *entity.Comment) error {
	now := time.Now()
	comment.CreatedAt = now
	comment.UpdatedAt = now
	return u.repo.Create(ctx, comment)
}

func (u *commentUsecaseImpl) UpdateComment(ctx context.Context, comment *entity.Comment) error {
	comment.UpdatedAt = time.Now()
	return u.repo.Update(ctx, comment)
}

func (u *commentUsecaseImpl) DeleteComment(ctx context.Context, commentID string) error {
	return u.repo.Delete(ctx, commentID)
}

func (u *commentUsecaseImpl) ListCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error) {
	return u.repo.ListByPostID(ctx, postID)
}
