package grpc

import (
	"context"

	pb "github.com/KaminurOrynbek/BiznesAsh/UserService/auto-proto/user"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/usecase/Usecase_Interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	userUsecase Usecase_Interfaces.UserUsecase
}

func NewUserServer(userUsecase Usecase_Interfaces.UserUsecase) *UserServer {
	return &UserServer{userUsecase: userUsecase}
}

func (s *UserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := s.userUsecase.Register(req.Email, req.Username, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register: %v", err)
	}

	// После успешной регистрации сразу выполняем вход для генерации токена
	_, token, err := s.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token after registration: %v", err)
	}

	return &pb.RegisterResponse{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
		Token:    token,
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, token, err := s.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to login: %v", err)
	}

	return &pb.LoginResponse{
		UserId: user.ID,
		Token:  token,
	}, nil
}

func (s *UserServer) Authorize(ctx context.Context, req *pb.TokenRequest) (*pb.AuthorizationResponse, error) {
	success, message, err := s.userUsecase.Authorize(req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to authorize: %v", err)
	}

	// Extract user_id from context (set by AuthInterceptor)
	userID, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to get user_id from context")
	}

	return &pb.AuthorizationResponse{
		Success: success,
		Message: message,
		UserId:  userID, // Populate the new field
	}, nil
}

func (s *UserServer) GetCurrentUser(ctx context.Context, _ *pb.Empty) (*pb.UserResponse, error) {
	userID, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}

	user, err := s.userUsecase.GetCurrentUser(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get current user: %v", err)
	}

	return &pb.UserResponse{
		UserId:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     string(user.Role),
		Bio:      user.Bio,
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := s.userUsecase.GetUser(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	return &pb.UserResponse{
		UserId:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     string(user.Role),
		Bio:      user.Bio,
	}, nil
}

func (s *UserServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UserResponse, error) {
	userID, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}

	user, err := s.userUsecase.UpdateProfile(userID, req.Username, req.Bio)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update profile: %v", err)
	}

	return &pb.UserResponse{
		UserId:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     string(user.Role),
		Bio:      user.Bio,
	}, nil
}

func (s *UserServer) PromoteToModerator(ctx context.Context, req *pb.RoleChangeRequest) (*pb.RoleChangeResponse, error) {
	currentUserID, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}

	_, err := s.userUsecase.PromoteToModerator(currentUserID, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to promote to moderator: %v", err)
	}

	return &pb.RoleChangeResponse{
		Success: true,
		Message: "User promoted to moderator",
	}, nil
}

func (s *UserServer) PromoteToAdmin(ctx context.Context, req *pb.RoleChangeRequest) (*pb.RoleChangeResponse, error) {
	currentUserID, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}

	_, err := s.userUsecase.PromoteToAdmin(currentUserID, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to promote to admin: %v", err)
	}

	return &pb.RoleChangeResponse{
		Success: true,
		Message: "User promoted to admin",
	}, nil
}

func (s *UserServer) DemoteToUser(ctx context.Context, req *pb.RoleChangeRequest) (*pb.RoleChangeResponse, error) {
	currentUserID, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}

	_, err := s.userUsecase.DemoteToUser(currentUserID, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to demote to user: %v", err)
	}

	return &pb.RoleChangeResponse{
		Success: true,
		Message: "User demoted to user",
	}, nil
}

func (s *UserServer) DeleteAccount(ctx context.Context, req *pb.UserID) (*pb.DeleteResponse, error) {
	currentUserID, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}

	err := s.userUsecase.DeleteAccount(currentUserID, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete account: %v", err)
	}

	return &pb.DeleteResponse{
		Success: true,
		Message: "Account deleted successfully",
	}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UsersListResponse, error) {
	users, err := s.userUsecase.ListUsers(req.SearchQuery)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users: %v", err)
	}

	response := &pb.UsersListResponse{}
	for _, u := range users {
		response.Users = append(response.Users, &pb.UserResponse{
			UserId:   u.ID,
			Email:    u.Email,
			Username: u.Username,
			Role:     string(u.Role),
			Bio:      u.Bio,
		})
	}

	return response, nil
}

func (s *UserServer) BanUser(ctx context.Context, req *pb.UserID) (*pb.BanUserResponse, error) {
	currentUserID, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}

	err := s.userUsecase.BanUser(currentUserID, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to ban user: %v", err)
	}

	return &pb.BanUserResponse{
		Success: true,
		Message: "User banned successfully",
	}, nil
}
