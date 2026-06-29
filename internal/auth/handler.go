package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Call business logic
	resp, err := h.service.Register(c.Request.Context(), req)

	if err != nil {
		slog.Error("Register failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Create user:", "id", resp.ID, "full_name", resp.FullName, "email", resp.Email, "created_at", resp.CreatedAt)
	// Return response
	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c *gin.Context){
	var req LoginRequest

	// Parse JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Info("JSON parse user error:", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Call business logic
	resp, err := h.service.Login(c.Request.Context(), req)

	if err != nil {
		slog.Error("Login failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Create user:", "id", resp.ID, "full_name", resp.FullName, "email", resp.Email, "created_at", resp.CreatedAt)
	// Return response
	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID := c.GetString("user_id")

	resp, err := h.service.GetUserbyID(c.Request.Context(), userID)

	if err != nil {
		slog.Error("Me failed", "user_id", userID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Me", "user_id", resp.ID, "email", resp.Email)
	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) GetUsers(c * gin.Context){
	resps, err := h.service.GetUsers(c.Request.Context())

	if err != nil {
		slog.Error("Get users failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Users", "users", resps)
	c.JSON(http.StatusOK, resps)
}