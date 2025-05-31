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

func (dao *LikeDAO) Like(ctx context.Context, like *model.Like) (int32, error) {
	query := `
		INSERT INTO likes (id, post_id, user_id, is_like, created_at)
		VALUES (:id, :post_id, :user_id, :is_like, :created_at)
		ON CONFLICT (post_id, user_id) DO NOTHING;
	`
	_, err := dao.db.NamedExecContext(ctx, query, like)
	if err != nil {
		return 0, err
	}
	return dao.CountLikes(ctx, like.PostID.String())
}

func (dao *LikeDAO) Dislike(ctx context.Context, dislike *model.Like) (int32, error) {
	query := `
		INSERT INTO likes (id, post_id, user_id, is_like, created_at)
		VALUES (:id, :post_id, :user_id, :is_like, :created_at)
		ON CONFLICT (post_id, user_id) DO NOTHING;
	`

	_, err := dao.db.NamedExecContext(ctx, query, dislike)
	if err != nil {
		return 0, err
	}

	return dao.CountDislikes(ctx, dislike.PostID.String())
}

func (dao *LikeDAO) CountLikes(ctx context.Context, postID string) (int32, error) {
	query := `
		SELECT COUNT(*) FROM likes 
		WHERE post_id = $1 AND is_like = true
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

func (dao *LikeDAO) LikeComment(ctx context.Context, like *model.Like) (int32, error) {
	query := `
		INSERT INTO likes (id, comment_id, user_id, is_like, created_at)
		VALUES (:id, :comment_id, :user_id, :is_like, :created_at)
		ON CONFLICT (comment_id, user_id) DO NOTHING;
	`
	_, err := dao.db.NamedExecContext(ctx, query, like)
	if err != nil {
		return 0, err
	}
	return dao.CountCommentLikes(ctx, like.CommentID.String())
}

func (dao *LikeDAO) CountCommentLikes(ctx context.Context, commentID string) (int32, error) {
	query := `SELECT COUNT(*) FROM likes WHERE comment_id = $1 AND is_like = true`
	var count int32
	err := dao.db.GetContext(ctx, &count, query, commentID)
	return count, err
}
