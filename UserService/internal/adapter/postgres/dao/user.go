package dao

import (
	"context"
	"fmt"
	"github.com/KaminurOrynbek/BiznesAsh/internal/repository/RepoInterfaces"

	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dto"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserDAO struct {
	db *sqlx.DB
}

func NewUserDAO(db *sqlx.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (d *UserDAO) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	dtoUser := dto.ToUserDTO(user)
	query := `
        INSERT INTO users (id, email, username, password, role, bio, banned, created_at, updated_at)
        VALUES (:id, :email, :username, :password, :role, :bio, :banned, :created_at, :updated_at)
    `
	_, err := d.db.NamedExecContext(ctx, query, dtoUser)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // Unique violation
			if pqErr.Constraint == "users_email_key" {
				return nil, fmt.Errorf("email %s already exists", user.Email)
			}
			if pqErr.Constraint == "users_username_key" {
				return nil, fmt.Errorf("username %s already exists", user.Username)
			}
		}
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func (d *UserDAO) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	var dtoUser dto.UserDTO
	query := `SELECT * FROM users WHERE id = $1`
	err := d.db.GetContext(ctx, &dtoUser, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %v", err)
	}
	return dto.ToEntityUser(&dtoUser), nil
}

func (d *UserDAO) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var dtoUser dto.UserDTO
	query := `SELECT * FROM users WHERE email = $1`
	err := d.db.GetContext(ctx, &dtoUser, query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}
	return dto.ToEntityUser(&dtoUser), nil
}

func (d *UserDAO) UpdateUser(ctx context.Context, user *entity.User) error {
	dtoUser := dto.ToUserDTO(user)
	query := `
        UPDATE users
        SET email = :email, username = :username, password = :password, role = :role,
            bio = :bio, banned = :banned, updated_at = :updated_at
        WHERE id = :id
    `
	_, err := d.db.NamedExecContext(ctx, query, dtoUser)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (d *UserDAO) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := d.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func (d *UserDAO) ListUsers(ctx context.Context, filter RepoInterfaces.UserFilter) ([]*entity.User, error) {
	var dtoUsers []dto.UserDTO
	// Build the query based on the filter
	query := `SELECT * FROM users WHERE 1=1`

	// Use the filter fields to build the WHERE clause dynamically
	args := []interface{}{}
	if filter.Email != "" {
		query += ` AND email ILIKE $` + fmt.Sprint(len(args)+1)
		args = append(args, "%"+filter.Email+"%")
	}
	if filter.Username != "" {
		query += ` AND username ILIKE $` + fmt.Sprint(len(args)+1)
		args = append(args, "%"+filter.Username+"%")
	}
	if filter.Role != "" {
		query += ` AND role = $` + fmt.Sprint(len(args)+1)
		args = append(args, filter.Role)
	}
	if filter.Banned != nil {
		query += ` AND banned = $` + fmt.Sprint(len(args)+1)
		args = append(args, *filter.Banned)
	}

	// Add limit and offset if they are set
	if filter.Limit > 0 {
		query += ` LIMIT $` + fmt.Sprint(len(args)+1)
		args = append(args, filter.Limit)
	}
	if filter.Offset > 0 {
		query += ` OFFSET $` + fmt.Sprint(len(args)+1)
		args = append(args, filter.Offset)
	}

	// Execute the query
	err := d.db.SelectContext(ctx, &dtoUsers, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %v", err)
	}

	// Convert DTOs to entities
	users := make([]*entity.User, len(dtoUsers))
	for i, dtoUser := range dtoUsers {
		users[i] = dto.ToEntityUser(&dtoUser)
	}
	return users, nil
}

func (d *UserDAO) BanUser(ctx context.Context, id string) error {
	query := `UPDATE users SET banned = true, updated_at = CURRENT_TIMESTAMP WHERE id = $1`
	_, err := d.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to ban user: %v", err)
	}
	return nil
}
