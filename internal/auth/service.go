package auth

import (
	"context"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {

	user, err := s.repo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.HashPassword),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := GenerateAccessToken(user)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: token,
		Role:        string(user.Role),
	}, nil
}

func (s *Service) Me(ctx context.Context, userID uint) (*MeResponse, error) {
	user, err := s.repo.GetUserByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return &MeResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	// Generate password
	password := GeneratePassword(16)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &database.User{
		Name:         req.Name,
		Email:        req.Email,
		Role:         database.UserRole(req.Role),
		HashPassword: string(hashedPassword),
		IsActive:     true,
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     string(user.Role),
		Password: password,
	}, nil
}
