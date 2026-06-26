package main

import (
	"log/slog"
	"net/http"
	"os"
	"github.com/Khoa-BOB/hireflow/internal/logger"
	"github.com/Khoa-BOB/hireflow/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	log := logger.NewLogger()
	slog.SetDefault(log)
	slog.Info("Starting server on localhost:8000")

	router.GET("/health", func(c *gin.Context) {
		conn, err := database.NewPostgresConnection()
		if err != nil {
			slog.Error("Database connection failed", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "ERROR", "database": "Connection failed"})
			return
		}
		defer conn.Close()
		slog.Info("Connected to PostgreSQL")
		c.IndentedJSON(http.StatusOK, gin.H{"status": "OK", "database": "Connected"})
	})
	
	if err := router.Run("localhost:8000"); err != nil {
		slog.Error("Failed to start the server", "error", err)
		os.Exit(1)
	}
	slog.Info("Server stopped")
}
