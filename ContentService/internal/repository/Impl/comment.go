package Impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/jmoiron/sqlx"
)

type commentRepositoryImpl struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *commentRepositoryImpl {
	return &commentRepositoryImpl{db: db}
}

func (r *commentRepositoryImpl) Create(ctx context.Context, comment *entity.Comment) error {
	query := `INSERT INTO comments (id, post_id, author_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, comment.ID, comment.PostID, comment.AuthorID, comment.Content, comment.CreatedAt, comment.UpdatedAt)
	return err
}

func (r *commentRepositoryImpl) Update(ctx context.Context, comment *entity.Comment) error {
	query := `UPDATE comments SET content = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, comment.Content, comment.UpdatedAt, comment.ID)
	return err
}

func (r *commentRepositoryImpl) Delete(ctx context.Context, commentID string) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, commentID)
	return err
}

func (r *commentRepositoryImpl) ListByPostID(ctx context.Context, postID string) ([]*entity.Comment, error) {
	query := `SELECT id, post_id, author_id, content, created_at, updated_at FROM comments WHERE post_id = $1`
	var comments []*entity.Comment
	err := r.db.SelectContext(ctx, &comments, query, postID)
	return comments, err
}
