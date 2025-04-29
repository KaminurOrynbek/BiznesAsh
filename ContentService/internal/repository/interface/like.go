package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type LikeRepository interface {
	Like(ctx context.Context, like *entity.Like) error
	Dislike(ctx context.Context, like *entity.Like) error
	CountLikes(ctx context.Context, postID string) (int32, error)
	CountDislikes(ctx context.Context, postID string) (int32, error)
}
