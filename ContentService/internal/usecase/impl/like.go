package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	ucase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/google/uuid"
	"time"
)

type likeUsecaseImpl struct {
	repo _interface.LikeRepository
}

func NewLikeUsecase(repo _interface.LikeRepository) ucase.LikeUsecase {
	return &likeUsecaseImpl{repo: repo}
}

func (u *likeUsecaseImpl) LikePost(ctx context.Context, postID string, userID string) error {
	like := &entity.Like{
		ID:        uuid.NewString(),
		PostID:    postID,
		UserID:    userID,
		IsLike:    true,
		CreatedAt: time.Now(),
	}
	return u.repo.Like(ctx, like)
}

func (u *likeUsecaseImpl) DislikePost(ctx context.Context, postID string, userID string) error {
	like := &entity.Like{
		ID:        uuid.NewString(),
		PostID:    postID,
		UserID:    userID,
		IsLike:    false,
		CreatedAt: time.Now(),
	}
	return u.repo.Dislike(ctx, like)
}
