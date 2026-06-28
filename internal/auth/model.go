package auth

import "time"

type User struct {
	ID string
	FullName string
	Email string
	PasswordHash string
	Role string
	CreatedAt time.Time
}