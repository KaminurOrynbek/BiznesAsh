package repository

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/repository/RepoInterfaces"
)

type userRepositoryImpl struct {
	userDAO *dao.UserDAO
}

func NewUserRepository(userDAO *dao.UserDAO) RepoInterfaces.UserRepository {
	return &userRepositoryImpl{
		userDAO: userDAO,
	}
}

// Convert entity.UserFilter to RepoInterfaces.UserFilter
func convertToRepoInterfacesFilter(filter entity.UserFilter) RepoInterfaces.UserFilter {
	return RepoInterfaces.UserFilter{
		Email:    filter.Email,
		Username: filter.Username,
		Role:     filter.Role,
		Banned:   filter.Banned,
		Limit:    filter.Limit,
		Offset:   filter.Offset,
	}
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return r.userDAO.CreateUser(ctx, user)
}

func (r *userRepositoryImpl) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return r.userDAO.GetUserByID(ctx, id)
}

func (r *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return r.userDAO.GetUserByEmail(ctx, email)
}

func (r *userRepositoryImpl) UpdateUser(ctx context.Context, user *entity.User) error {
	return r.userDAO.UpdateUser(ctx, user)
}

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	return r.userDAO.DeleteUser(ctx, id)
}

func (r *userRepositoryImpl) ListUsers(ctx context.Context, filter entity.UserFilter) ([]*entity.User, error) {
	return r.userDAO.ListUsers(ctx, filter)
}

func (r *userRepositoryImpl) BanUser(ctx context.Context, id string) error {
	return r.userDAO.BanUser(ctx, id)
}
