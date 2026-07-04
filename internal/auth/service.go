package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service{
	return &Service{
		repo: repo,
	}
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error){

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

func (s *Service) Me(ctx context.Context, userID uint) (*MeResponse, error){
	user, err := s.repo.GetUserByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return &MeResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: string(user.Role),
	}, nil
}