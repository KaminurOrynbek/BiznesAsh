package Usecase_Interfaces

import (
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/entity"
)

type UserUsecase interface {
	Register(email, username, password string) (*entity.User, error)
	Login(email, password string) (*entity.User, string, error)
	Authorize(token string) (bool, string, error)
	GetCurrentUser(userId string) (*entity.User, error)
	GetUser(userId string) (*entity.User, error)
	UpdateProfile(userId, username, bio string) (*entity.User, error)
	PromoteToModerator(currentUserId, targetUserId string) (*entity.User, error)
	PromoteToAdmin(currentUserId, targetUserId string) (*entity.User, error)
	DemoteToUser(currentUserId, targetUserId string) (*entity.User, error)
	DeleteAccount(currentUserId, targetUserId string) error
	ListUsers(searchQuery string) ([]*entity.User, error)
	BanUser(currentUserId, targetUserId string) error
}
