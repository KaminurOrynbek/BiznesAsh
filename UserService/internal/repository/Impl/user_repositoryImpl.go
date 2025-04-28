package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/repository/RepoInterfaces"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) RepoInterfaces.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users (id, email, username, password, role, bio, banned, created_at, updated_at)
		VALUES (:id, :email, :username, :password, :role, :bio, :banned, :created_at, :updated_at)
		RETURNING *
	`
	var result entity.User
	namedQuery, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare query")
	}
	err = namedQuery.GetContext(ctx, &result, user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return &result, nil
}

func (r *userRepositoryImpl) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, &user, query, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, &user, query, email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by email")
	}
	return &user, nil
}

func (r *userRepositoryImpl) UpdateUser(ctx context.Context, user *entity.User) error {
	user.UpdatedAt = time.Now()
	query := `
		UPDATE users
		SET email = :email, username = :username, password = :password, role = :role,
		    bio = :bio, banned = :banned, updated_at = :updated_at
		WHERE id = :id
	`
	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to check affected rows")
	}
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *userRepositoryImpl) ListUsers(ctx context.Context, filter RepoInterfaces.UserFilter) ([]*entity.User, error) {
	query := `SELECT * FROM users WHERE 1=1`
	args := []interface{}{}
	if filter.Email != "" {
		query += ` AND email ILIKE $1`
		args = append(args, "%"+filter.Email+"%")
	}
	if filter.Username != "" {
		query += ` AND username ILIKE $2`
		args = append(args, "%"+filter.Username+"%")
	}
	if filter.Role != "" {
		query += ` AND role = $3`
		args = append(args, filter.Role)
	}
	if filter.Banned != nil {
		query += ` AND banned = $4`
		args = append(args, *filter.Banned)
	}
	query += ` LIMIT $5 OFFSET $6`
	args = append(args, filter.Limit, filter.Offset)

	var users []*entity.User
	err := r.db.SelectContext(ctx, &users, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}
	return users, nil
}
