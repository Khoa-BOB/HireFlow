package auth

import (
		"context"
		"time"
)

// Interface

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, uuid string) (*User, error)
	GetUsers(ctx context.Context) ([]UserWithRole, error)
	CreateUser(ctx context.Context, email, passwordHash, fullName string) (*User, error)
	CreateRefreshToken(ctx context.Context, id, refreshToken string, expiresAt time.Time) (*RefreshToken, error)
	GetRefreshTokenByHash(ctx context.Context, refreshTokenHash string) (*RefreshToken, error)
	RevokeRefreshToken(ctx context.Context, refreshTokenHash string) error
}