package auth

import (
	"context"
	"errors"
	"strings"
	"log/slog"
	"golang.org/x/crypto/bcrypt"
)

var errEmailAlreadyExists = errors.New("email already exists")

type AuthService struct {
	repo UserRepository
}

func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s * AuthService) Register(ctx context.Context, req RegisterRequest) (* RegisterResponse, error){
	email:= strings.ToLower(strings.TrimSpace(req.Email))
	fullName:= strings.TrimSpace(req.FullName)

	existingUser, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errEmailAlreadyExists
	}

	slog.Info("Hashing password", "full_name", fullName, "email", email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(ctx, email, string(hashedPassword), fullName)

	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}, nil

}