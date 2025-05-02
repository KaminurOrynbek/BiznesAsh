package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	usecase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"time"
)

type likeUsecaseImpl struct {
	likeRepo _interface.LikeRepository
}

func NewLikeUsecase(likeRepo _interface.LikeRepository) usecase.LikeUsecase {
	return &likeUsecaseImpl{likeRepo: likeRepo}
}

func (u *likeUsecaseImpl) LikePost(ctx context.Context, like *entity.Like) error {
	like.IsLike = true
	like.CreatedAt = time.Now()
	return u.likeRepo.Like(ctx, like)
}

func (u *likeUsecaseImpl) DislikePost(ctx context.Context, like *entity.Like) error {
	like.IsLike = false
	like.CreatedAt = time.Now()
	return u.likeRepo.Dislike(ctx, like)
}
