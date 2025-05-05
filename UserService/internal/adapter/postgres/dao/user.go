package dao

import (
	"context"
	"fmt"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/postgres/model"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/entity"
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
	dtoUser := model.ToUserDB(user)
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
	var dtoUser model.UserDB
	query := `SELECT * FROM users WHERE id = $1`
	err := d.db.GetContext(ctx, &dtoUser, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %v", err)
	}
	return model.ToEntityUser(&dtoUser), nil
}

func (d *UserDAO) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var dtoUser model.UserDB
	query := `SELECT * FROM users WHERE email = $1`
	err := d.db.GetContext(ctx, &dtoUser, query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}
	return model.ToEntityUser(&dtoUser), nil
}

func (d *UserDAO) UpdateUser(ctx context.Context, user *entity.User) error {
	dtoUser := model.ToUserDB(user)
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
func (d *UserDAO) ListUsers(ctx context.Context, filter entity.UserFilter) ([]*entity.User, error) {
	var dtoUsers []model.UserDB
	// Convert RepoInterfaces.UserFilter to entity.UserFilter
	entityFilter := entity.UserFilter{
		Email:    filter.Email,
		Username: filter.Username,
		Role:     filter.Role,
		Banned:   filter.Banned,
		Limit:    filter.Limit,
		Offset:   filter.Offset,
	}

	// Build the query using entityFilter
	query := `SELECT * FROM users WHERE 1=1`
	args := []interface{}{}
	if entityFilter.Email != "" {
		query += ` AND email ILIKE $` + fmt.Sprint(len(args)+1)
		args = append(args, "%"+entityFilter.Email+"%")
	}
	// (Other conditions here)

	// Execute the query
	err := d.db.SelectContext(ctx, &dtoUsers, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %v", err)
	}

	// Convert DTOs to entities
	users := make([]*entity.User, len(dtoUsers))
	for i, dtoUser := range dtoUsers {
		users[i] = model.ToEntityUser(&dtoUser)
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
