package company

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *CompanyService
}

func NewHandler(service *CompanyService) *CompanyHandler {
	return &CompanyHandler{
		service: service,
	}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var req CreateCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Info("JSON parse create company error:", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.service.CreateCompany(c.Request.Context(), &req)
	if err != nil {
		slog.Error("Create company failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Create company:", "company", resp)
	c.JSON(http.StatusCreated, resp)
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	id := c.Param("id")

	var req UpdateCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Info("JSON parse update company error:", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.service.UpdateCompany(c.Request.Context(), id, &req)
	if err != nil {
		slog.Error("Update company failed", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Update company:", "company", resp)
	c.JSON(http.StatusOK, resp)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteCompany(c.Request.Context(), id); err != nil {
		slog.Error("Delete company failed", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Delete company success", "id", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Company deleted successfully",
	})
}

func (h *CompanyHandler) GetCompanyByID(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.service.GetCompanyByID(c.Request.Context(), id)
	if err != nil {
		slog.Error("Get company by ID failed", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Get company by ID:", "company", resp)
	c.JSON(http.StatusOK, resp)
}

func (h *CompanyHandler) GetAllCompanies(c *gin.Context) {
	resps, err := h.service.GetAllCompanies(c.Request.Context())
	if err != nil {
		slog.Error("Get all companies failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	slog.Info("Get all companies", "count", len(resps))
	c.JSON(http.StatusOK, resps)
}
