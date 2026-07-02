package company

import "github.com/gin-gonic/gin"

func RegisterModule(router *gin.Engine, repo CompanyRepository) {
	service := NewCompanyService(repo)
	handler := NewHandler(service)

	group := router.Group("/companies")
	handler.RegisterRoutes(group)
}
