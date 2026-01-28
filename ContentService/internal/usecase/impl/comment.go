package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	usecase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/google/uuid"
	"time"
)

type commentUsecaseImpl struct {
	commentRepo _interface.CommentRepository
}

func NewCommentUsecase(commentRepo _interface.CommentRepository) usecase.CommentUsecase {
	return &commentUsecaseImpl{commentRepo: commentRepo}
}

func (u *commentUsecaseImpl) CreateComment(ctx context.Context, comment *entity.Comment) error {
	comment.ID = uuid.NewString()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = comment.CreatedAt
	return u.commentRepo.Create(ctx, comment)
}

func (u *commentUsecaseImpl) UpdateComment(ctx context.Context, comment *entity.Comment) error {
	comment.UpdatedAt = time.Now()
	return u.commentRepo.Update(ctx, comment)
}

func (u *commentUsecaseImpl) DeleteComment(ctx context.Context, commentID string) error {
	return u.commentRepo.Delete(ctx, commentID)
}

func (u *commentUsecaseImpl) ListCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error) {
	return u.commentRepo.ListByPostID(ctx, postID)
}
