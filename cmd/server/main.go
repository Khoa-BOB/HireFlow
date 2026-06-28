package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Khoa-BOB/hireflow/internal/auth"
	"github.com/Khoa-BOB/hireflow/internal/database"
	"github.com/Khoa-BOB/hireflow/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	router := gin.Default()
	log := logger.NewLogger()
	slog.SetDefault(log)
	slog.Info("Starting server on localhost:8000")
	
	// Create database connection
	post_db, err := database.NewPostgresConnection()
	if err != nil {
		slog.Error("Database connection failed", "error", err)
		os.Exit(1)
	}
	defer post_db.Close()

	// Health route
	router.GET("/health", func(c *gin.Context) {
		if err := post_db.Ping(c.Request.Context()); err != nil {
			slog.Error("Database ping failed", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   "ERROR",
				"database": "Connection failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "OK",
			"database": "Connected",
		})
	})

	// Register auth module
	auth.RegisterModule(router, auth.NewPostgresRepository(post_db))

	// Start the server

	if err := router.Run("localhost:8000"); err != nil {
		slog.Error("Failed to start the server", "error", err)
		os.Exit(1)
	}
	slog.Info("Server stopped")
}
