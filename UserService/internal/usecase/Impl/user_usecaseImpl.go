package usecase

import (
	"context"
	"os"
	"time"

	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity/enum"
	"github.com/KaminurOrynbek/BiznesAsh/internal/repository/RepoInterfaces"
	"github.com/KaminurOrynbek/BiznesAsh/internal/usecase/Usecase_Interfaces"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type userUsecaseImpl struct {
	userRepo RepoInterfaces.UserRepository
}

func NewUserUsecase(userRepo RepoInterfaces.UserRepository) Usecase_Interfaces.UserUsecase {
	return &userUsecaseImpl{userRepo: userRepo}
}

func (u *userUsecaseImpl) Register(email, username, password string) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash password")
	}

	user := &entity.User{
		ID:        uuid.New().String(),
		Email:     email,
		Username:  username,
		Password:  string(hashedPassword),
		Role:      enum.RoleUser,
		Bio:       "",
		Banned:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := u.userRepo.CreateUser(context.Background(), user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return createdUser, nil
}

func (u *userUsecaseImpl) Login(email, password string) (*entity.User, string, error) {
	user, err := u.userRepo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, "", errors.Wrap(err, "user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid password")
	}

	token, err := generateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to generate token")
	}

	return user, token, nil
}

func (u *userUsecaseImpl) Authorize(token string) (bool, string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !parsedToken.Valid {
		return false, "invalid token", errors.Wrap(err, "failed to parse token")
	}
	return true, "token valid", nil
}

func (u *userUsecaseImpl) GetCurrentUser(userID string) (*entity.User, error) {
	user, err := u.userRepo.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current user")
	}
	return user, nil
}

func (u *userUsecaseImpl) GetUser(userID string) (*entity.User, error) {
	user, err := u.userRepo.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	return user, nil
}

func (u *userUsecaseImpl) UpdateProfile(userID, username, bio string) (*entity.User, error) {
	user, err := u.userRepo.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}

	user.Username = username
	user.Bio = bio

	err = u.userRepo.UpdateUser(context.Background(), user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update profile")
	}

	return user, nil
}

func (u *userUsecaseImpl) PromoteToModerator(currentUserID, targetUserID string) (*entity.User, error) {
	currentUser, err := u.userRepo.GetUserByID(context.Background(), currentUserID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current user")
	}
	if !currentUser.Role.IsAdmin() {
		return nil, errors.New("only admins can promote to moderator")
	}

	targetUser, err := u.userRepo.GetUserByID(context.Background(), targetUserID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get target user")
	}

	targetUser.Role = enum.RoleModerator
	err = u.userRepo.UpdateUser(context.Background(), targetUser)
	if err != nil {
		return nil, errors.Wrap(err, "failed to promote to moderator")
	}

	return targetUser, nil
}

func (u *userUsecaseImpl) PromoteToAdmin(currentUserID, targetUserID string) (*entity.User, error) {
	currentUser, err := u.userRepo.GetUserByID(context.Background(), currentUserID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current user")
	}
	if !currentUser.Role.IsAdmin() {
		return nil, errors.New("only admins can promote to admin")
	}

	targetUser, err := u.userRepo.GetUserByID(context.Background(), targetUserID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get target user")
	}

	targetUser.Role = enum.RoleAdmin
	err = u.userRepo.UpdateUser(context.Background(), targetUser)
	if err != nil {
		return nil, errors.Wrap(err, "failed to promote to admin")
	}

	return targetUser, nil
}

func (u *userUsecaseImpl) DemoteToUser(currentUserID, targetUserID string) (*entity.User, error) {
	currentUser, err := u.userRepo.GetUserByID(context.Background(), currentUserID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current user")
	}
	if !currentUser.Role.IsAdmin() {
		return nil, errors.New("only admins can demote to user")
	}

	targetUser, err := u.userRepo.GetUserByID(context.Background(), targetUserID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get target user")
	}

	targetUser.Role = enum.RoleUser
	err = u.userRepo.UpdateUser(context.Background(), targetUser)
	if err != nil {
		return nil, errors.Wrap(err, "failed to demote to user")
	}

	return targetUser, nil
}

func (u *userUsecaseImpl) DeleteAccount(currentUserID, targetUserID string) error {
	currentUser, err := u.userRepo.GetUserByID(context.Background(), currentUserID)
	if err != nil {
		return errors.Wrap(err, "failed to get current user")
	}
	if !currentUser.Role.IsAdmin() {
		return errors.New("only admins can delete accounts")
	}

	err = u.userRepo.DeleteUser(context.Background(), targetUserID)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

func (u *userUsecaseImpl) ListUsers(searchQuery string) ([]*entity.User, error) {
	filter := entity.UserFilter{
		SearchQuery: searchQuery,
	}

	users, err := u.userRepo.ListUsers(context.Background(), filter)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}
	return users, nil
}

// BanUser sets the 'Banned' status of a user to true.
func (u *userUsecaseImpl) BanUser(currentUserId string, targetUserId string) error {
	currentUser, err := u.userRepo.GetUserByID(context.Background(), currentUserId)
	if err != nil {
		return errors.Wrap(err, "failed to get current user")
	}
	if !currentUser.Role.IsAdmin() {
		return errors.New("only admins can ban users")
	}

	targetUser, err := u.userRepo.GetUserByID(context.Background(), targetUserId)
	if err != nil {
		return errors.Wrap(err, "failed to get target user")
	}

	targetUser.Banned = true
	err = u.userRepo.UpdateUser(context.Background(), targetUser)
	if err != nil {
		return errors.Wrap(err, "failed to ban user")
	}

	return nil
}

func generateToken(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token")
	}
	return tokenString, nil
}
