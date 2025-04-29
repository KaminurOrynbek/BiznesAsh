package _interface

import "context"

type LikeUsecase interface {
	LikePost(ctx context.Context, postID, userID string) error
	DislikePost(ctx context.Context, postID, userID string) error
}
