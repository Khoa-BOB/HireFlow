package company

import (
	"context"
)

// Interface
type CompanyRepository interface {
	CreateCompany(ctx context.Context, company *Company) (*Company, error)
	UpdateCompany(ctx context.Context, id string, company *CompanyUpdate) (*Company, error)
	DeleteCompany(ctx context.Context, id string) error
	GetCompanyByID(ctx context.Context, id string) (*Company, error)
	GetAllCompanies(ctx context.Context) ([]*Company, error)
}