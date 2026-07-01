package auth

import (
	"context"
	"errors"
	"strings"
	"log/slog"
	"golang.org/x/crypto/bcrypt"
	"time"
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
		Role: 	   user.Role,
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

	slog.Info("Generate JWT token...")
	access_token, err := GenerateJWT(user.ID, user.Email, user.Role)

	if err != nil {
		return nil, err
	}

	slog.Info("Generate Refresh token...")
	refresh_token, err := GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	refresh_token_hash := HashRefreshToken(refresh_token)
	
	_, err = s.repo.CreateRefreshToken(ctx, user.ID, refresh_token_hash, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: 	access_token,
		RefreshToken: refresh_token,
		User: UserResponse{
			ID:        		user.ID,
			Email:     		user.Email,
			FullName:  		user.FullName,
			Role: 			user.Role,
			CreatedAt: 		user.CreatedAt,
		},
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	refreshTokenHash := HashRefreshToken(refreshToken)
	
	token, err := s.repo.GetRefreshTokenByHash(ctx, refreshTokenHash)
	if err != nil {
		return err
	}
	
	if token == nil || token.RevokedAt != nil || token.ExpiresAt.Before(time.Now()) {
		return nil // already revoked/expired/unknown - nothing to do, logout succeeds
	}

	err = s.repo.RevokeRefreshToken(ctx, refreshTokenHash)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GetUserbyID(ctx context.Context, id string) (*UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errUserNotFound
	}

	return &UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *AuthService) GetUsers(ctx context.Context) ([]UserResponse, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	resps := make([]UserResponse, 0, len(users))
	for _, user := range users {
		resps = append(resps, UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FullName:  user.FullName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		})
	}

	return resps, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*LoginResponse, error) {
	refreshTokenHash := HashRefreshToken(refreshToken)
	
	token, err := s.repo.GetRefreshTokenByHash(ctx, refreshTokenHash)
	if err != nil {
		return nil, err
	}

	if token == nil || token.RevokedAt != nil || token.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("invalid or expired refresh token")
	}

	user, err := s.repo.GetUserByID(ctx, token.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errUserNotFound
	}

	slog.Info("Generate new JWT token...")
	newAccessToken, err := GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	slog.Info("Generate new Refresh token...")
	newRefreshToken, err := GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	newRefreshTokenHash := HashRefreshToken(newRefreshToken)

	// Revoke the old refresh token
	err = s.repo.RevokeRefreshToken(ctx, refreshTokenHash)
	if err != nil {
		return nil, err
	}

	// Store the new refresh token
	_, err = s.repo.CreateRefreshToken(ctx, user.ID, newRefreshTokenHash, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		User: UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FullName:  user.FullName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}