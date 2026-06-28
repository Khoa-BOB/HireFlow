package auth

import (
	"context"
	"errors"
	"strings"
	"log/slog"
	"golang.org/x/crypto/bcrypt"
)

var errEmailAlreadyExists = errors.New("email already exists")
var errInvalidCredentials = errors.New("invalid email or password")
var errUserNotFound = errors.New("user not found") 
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

func (s * AuthService) Login(ctx context.Context, req LoginRequest) (* LoginResponse, error){
	email := strings.ToLower(strings.TrimSpace(req.Email))

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errInvalidCredentials
	}

	// Compare the provided password with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errInvalidCredentials
	}

	slog.Info("Generate JWT token")
	token, err := GenerateJWT(user.ID, user.Email)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: 	token,
		ID:        		user.ID,
		Email:     		user.Email,
		FullName:  		user.FullName,
		CreatedAt: 		user.CreatedAt,
	}, nil
}

func (s * AuthService) GetUserbyID(ctx context.Context, id string) (* UserResponse, error){
	
	user, err := s.repo.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	
	if user == nil {
		return nil, errUserNotFound
	}

	return &UserResponse{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
	}, nil
}