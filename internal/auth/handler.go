package auth

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type AuthHandler struct {
	service *AuthService
}

func NewHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// Parse JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Info("JSON parse user error:", "error", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Call business logic
	resp, err := h.service.Register(c.Request.Context(), req)
	
	if err != nil {
		slog.Info("Server error:", "error", err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Create user:", "id", resp.ID, "full_name", resp.FullName, "email", resp.Email, "created_at", resp.CreatedAt)
	// Return response
	c.JSON(201, resp)
}