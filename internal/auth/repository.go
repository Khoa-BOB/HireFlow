package auth

import (
		"context"
)

// Interface

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, uuid string) (*User, error)
	CreateUser(ctx context.Context, email, passwordHash, fullName string) (*User, error)
}