package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()

	router.GET("/", func(c * gin.Context){
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello from server!"})
	})
	router.Run("localhost:8000")
}