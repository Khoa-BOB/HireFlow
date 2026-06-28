package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var errMissingAuthHeader = errors.New("missing authorization header")
var errInvalidAuthHeader = errors.New("invalid authorization header")
var errInvalidToken = errors.New("invalid or expired token")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			slog.Warn("Missing authorization header", "path", c.Request.URL.Path)
			c.JSON(http.StatusUnauthorized, gin.H{"error": errMissingAuthHeader.Error()})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			slog.Warn("Invalid authorization header format", "path", c.Request.URL.Path)
			c.JSON(http.StatusUnauthorized, gin.H{"error": errInvalidAuthHeader.Error()})
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims := &JWTClaims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			slog.Warn("Invalid or expired token", "path", c.Request.URL.Path, "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": errInvalidToken.Error()})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		
		c.Next()
	}
}