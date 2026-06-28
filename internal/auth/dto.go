package auth

import "time"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ID        	string    `json:"id"`
	Email     	string    `json:"email"`
	FullName  	string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserResponse struct {
	ID string
	FullName string
	Email string
}
