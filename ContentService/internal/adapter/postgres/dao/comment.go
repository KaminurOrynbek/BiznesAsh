package dao

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/jmoiron/sqlx"
)

type CommentDAO struct {
	db *sqlx.DB
}

func NewCommentDAO(db *sqlx.DB) *CommentDAO {
	return &CommentDAO{db: db}
}

func (dao *CommentDAO) Create(ctx context.Context, comment *model.Comment) error {
	query := `
		INSERT INTO comments (id, post_id, author_id, content, created_at, updated_at)
		VALUES (:id, :post_id, :author_id, :content, :created_at, :updated_at)
	`
	_, err := dao.db.NamedExecContext(ctx, query, comment)
	return err
}

func (dao *CommentDAO) Update(ctx context.Context, comment *model.Comment) error {
	query := `
		UPDATE comments
		SET content = :content, updated_at = :updated_at
		WHERE id = :id
	`
	_, err := dao.db.NamedExecContext(ctx, query, comment)
	return err
}

func (dao *CommentDAO) Delete(ctx context.Context, commentID string) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := dao.db.ExecContext(ctx, query, commentID)
	return err
}

func (dao *CommentDAO) ListByPostID(ctx context.Context, postID string) ([]*model.Comment, error) {
	query := `
		SELECT id, post_id, author_id, content, created_at, updated_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at ASC
	`
	var comments []*model.Comment
	err := dao.db.SelectContext(ctx, &comments, query, postID)
	return comments, err
}
