package Impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/jmoiron/sqlx"
)

type likeRepositoryImpl struct {
	db *sqlx.DB
}

func NewLikeRepository(db *sqlx.DB) *likeRepositoryImpl {
	return &likeRepositoryImpl{db: db}
}

func (r *likeRepositoryImpl) Like(ctx context.Context, like *entity.Like) error {
	query := `INSERT INTO likes (id, post_id, user_id, is_like, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, like.ID, like.PostID, like.UserID, true, like.CreatedAt)
	return err
}

func (r *likeRepositoryImpl) Dislike(ctx context.Context, like *entity.Like) error {
	query := `INSERT INTO likes (id, post_id, user_id, is_like, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, like.ID, like.PostID, like.UserID, false, like.CreatedAt)
	return err
}

func (r *likeRepositoryImpl) CountLikes(ctx context.Context, postID string) (int32, error) {
	query := `SELECT COUNT(*) FROM likes WHERE post_id = $1 AND is_like = true`
	var count int32
	err := r.db.GetContext(ctx, &count, query, postID)
	return count, err
}

func (r *likeRepositoryImpl) CountDislikes(ctx context.Context, postID string) (int32, error) {
	query := `SELECT COUNT(*) FROM likes WHERE post_id = $1 AND is_like = false`
	var count int32
	err := r.db.GetContext(ctx, &count, query, postID)
	return count, err
}
