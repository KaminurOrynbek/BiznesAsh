package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type LikeUsecase interface {
	LikePost(ctx context.Context, like *entity.Like) (int32, error)
	DislikePost(ctx context.Context, like *entity.Like) (int32, error)
	LikeComment(ctx context.Context, like *entity.Like) (int32, error)
}
