package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Khoa-BOB/hireflow/internal/logger"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	log := logger.NewLogger()
	slog.SetDefault(log)
	slog.Info("Starting server on localhost:8000")
	router.GET("/", func(c * gin.Context){
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello from server!"})
	})
	if err:= router.Run("localhost:8000"); err != nil{
		slog.Error("Failed to start the server", "error", err)
		os.Exit(1)
	}
	slog.Info("Server stopped")
}