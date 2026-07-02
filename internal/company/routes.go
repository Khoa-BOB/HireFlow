package company

import (
	"github.com/gin-gonic/gin"
)

func (h *CompanyHandler) RegisterRoutes(router *gin.RouterGroup) {
	// public routes - no auth required
	public := router.Group("")
	public.GET("", h.GetAllCompanies)
	public.GET("/:id", h.GetCompanyByID)

	// admin-only routes
	admin := router.Group("")
	admin.Use(RequireAuth(), RequireAdmin())
	admin.POST("", h.CreateCompany)
	admin.PATCH("/:id", h.UpdateCompany)
	admin.DELETE("/:id", h.DeleteCompany)
}