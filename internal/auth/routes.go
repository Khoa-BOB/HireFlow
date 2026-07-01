package auth

import (
	"github.com/gin-gonic/gin"
)

func (h* AuthHandler) RegisterRoutes(router * gin.RouterGroup){
	//Public routes
	router.POST("/registry", h.Register)
	router.POST("/login",h.Login)
	router.POST("/refresh", h.RefreshToken)

	// protected routes
	protected := router.Group("")
	protected.Use(AuthMiddleware())
	protected.GET("/me", h.Me)

	admin := protected.Group("/admin")
	admin.Use(RequireRole("admin"))
	admin.GET("/users", h.GetUsers) 
}