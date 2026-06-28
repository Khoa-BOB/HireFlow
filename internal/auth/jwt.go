package auth

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID string, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := JWTClaims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	
	return token.SignedString([]byte(secret))
}