package company

import (
	"github.com/Khoa-BOB/hireflow/internal/auth"
	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return auth.AuthMiddleware()
}

func RequireAdmin() gin.HandlerFunc {
	return auth.RequireRole("admin")
}
