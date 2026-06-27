package auth

import "github.com/gin-gonic/gin"

func RegisterModule(router *gin.Engine, repo UserRepository){
	service := NewAuthService(repo)
	handler := NewHandler(service)

	group := router.Group("/auth")
	handler.RegisterRoutes(group)
}