package dao

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/jmoiron/sqlx"
)

type LikeDAO struct {
	db *sqlx.DB
}

func NewLikeDAO(db *sqlx.DB) *LikeDAO {
	return &LikeDAO{db: db}
}

func (dao *LikeDAO) Like(ctx context.Context, like *model.Like) error {
	query := `
		INSERT INTO likes (id, post_id, user_id, is_like, created_at)
		VALUES (:id, :post_id, :user_id, :is_like, :created_at)
	`
	_, err := dao.db.NamedExecContext(ctx, query, like)
	return err
}

func (dao *LikeDAO) Dislike(ctx context.Context, dislike *model.Like) error {
	query := `
		INSERT INTO likes (id, post_id, user_id, is_like, created_at)
		VALUES (:id, :post_id, :user_id, :is_like, :created_at)
	`
	_, err := dao.db.NamedExecContext(ctx, query, dislike)
	return err
}

func (dao *LikeDAO) CountLikes(ctx context.Context, postID string) (int32, error) {
	query := `
		SELECT COUNT(*) FROM likes WHERE post_id = $1 AND is_like = true
	`
	var count int32
	err := dao.db.GetContext(ctx, &count, query, postID)
	return count, err
}

func (dao *LikeDAO) CountDislikes(ctx context.Context, postID string) (int32, error) {
	query := `
		SELECT COUNT(*) FROM likes WHERE post_id = $1 AND is_like = false
	`
	var count int32
	err := dao.db.GetContext(ctx, &count, query, postID)
	return count, err
}
