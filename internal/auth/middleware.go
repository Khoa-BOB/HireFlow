package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

		claims, err := ValidateJWT(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": errInvalidToken.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		
		c.Next()
	}
}

func RequireRole(roles ...string) gin.HandlerFunc {
	allowed := map[string]bool{}
	for _, role := range roles {
		allowed[role] = true
	}
	return func(c * gin.Context){
		role := c.GetString("role")
		if role == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authentication",
			})
			return
		}
		if _, ok := allowed[role]; !ok{
			slog.Warn("Forbidden: insufficient role", "role", role, "path", c.Request.URL.Path)
			c.JSON(http.StatusForbidden, gin.H{
				"error" : "forbidden",
			})
			return
		}
		c.Next()
	}
}