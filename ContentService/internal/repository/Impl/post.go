package Impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/jmoiron/sqlx"
)

type postRepositoryImpl struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *postRepositoryImpl {
	return &postRepositoryImpl{db: db}
}

func (r *postRepositoryImpl) Create(ctx context.Context, post *entity.Post) error {
	query := `
		INSERT INTO posts (id, title, content, type, author_id, created_at, updated_at, published, likes_count, dislikes_count)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.ExecContext(ctx, query,
		post.ID,
		post.Title,
		post.Content,
		post.Type,
		post.AuthorID,
		post.CreatedAt,
		post.UpdatedAt,
		post.Published,
		post.LikesCount,
		post.DislikesCount,
	)
	return err
}

func (r *postRepositoryImpl) Update(ctx context.Context, post *entity.Post) error {
	query := `
		UPDATE posts 
		SET title = $1, content = $2, type = $3, updated_at = $4, published = $5
		WHERE id = $6
	`
	_, err := r.db.ExecContext(ctx, query,
		post.Title,
		post.Content,
		post.Type,
		post.UpdatedAt,
		post.Published,
		post.ID,
	)
	return err
}

func (r *postRepositoryImpl) Delete(ctx context.Context, postID string) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, postID)
	return err
}

func (r *postRepositoryImpl) GetByID(ctx context.Context, postID string) (*entity.Post, error) {
	query := `
		SELECT id, title, content, type, author_id, created_at, updated_at, published, likes_count, dislikes_count
		FROM posts
		WHERE id = $1
	`
	var post entity.Post
	err := r.db.GetContext(ctx, &post, query, postID)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepositoryImpl) List(ctx context.Context, offset, limit int) ([]*entity.Post, error) {
	query := `
		SELECT id, title, content, type, author_id, created_at, updated_at, published, likes_count, dislikes_count
		FROM posts
		ORDER BY created_at DESC
		OFFSET $1 LIMIT $2
	`
	var posts []*entity.Post
	err := r.db.SelectContext(ctx, &posts, query, offset, limit)
	return posts, err
}

func (r *postRepositoryImpl) Search(ctx context.Context, keyword string) ([]*entity.Post, error) {
	query := `
		SELECT id, title, content, type, author_id, created_at, updated_at, published, likes_count, dislikes_count
		FROM posts
		WHERE title ILIKE '%' || $1 || '%' OR content ILIKE '%' || $1 || '%'
		ORDER BY created_at DESC
		LIMIT 100
	`
	var posts []*entity.Post
	err := r.db.SelectContext(ctx, &posts, query, keyword)
	return posts, err
}
