package auth

import (
		"context"
)

// Interface

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, uuid string) (*User, error)
	GetUsers(ctx context.Context) ([]UserWithRole, error)
	CreateUser(ctx context.Context, email, passwordHash, fullName string) (*User, error)
}