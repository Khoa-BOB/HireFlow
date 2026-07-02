package company

import (
	"context"
	"log/slog"
	"time"
)

type CompanyService struct {
	repo CompanyRepository
}

func NewCompanyService(repo CompanyRepository) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

func (s *CompanyService) CreateCompany(ctx context.Context, req *CreateCompanyRequest) (*CompanyResponse, error) {
	comapny := &Company{
		Name:        req.Name,
		Description: req.Description,
		Industry:    req.Industry,
		URL:         req.URL,
		Address:     req.Address,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
		ZipCode:     req.ZipCode,
	}

	slog.Info("Creating company", "comapny", comapny)

	company, err := s.repo.CreateCompany(ctx, comapny)
	if err != nil {
		return nil, err
	}
		
	if err != nil {
		return nil, err
	}

	return toCompanyResponse(company), nil
}

func (s *CompanyService) UpdateCompany(ctx context.Context, id string, req *UpdateCompanyRequest) (*CompanyResponse, error) {
	updateData := &CompanyUpdate{
		Name:        req.Name,
		Description: req.Description,
		Industry:    req.Industry,
		URL:         req.URL,
		Address:     req.Address,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
		ZipCode:     req.ZipCode,
	}

	slog.Info("Updating company", "id", id, "updateData", updateData)

	company, err := s.repo.UpdateCompany(ctx, id, updateData)
	if err != nil {
		return nil, err
	}

	return toCompanyResponse(company), nil
}


func (s *CompanyService) DeleteCompany(ctx context.Context, id string) error {
	slog.Info("Deleting company", "id", id)
	return s.repo.DeleteCompany(ctx, id)
}

func (s *CompanyService) GetCompanyByID(ctx context.Context, id string) (*CompanyResponse, error) {
	slog.Info("Getting company by ID", "id", id)

	company, err := s.repo.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toCompanyResponse(company), nil
}

func (s *CompanyService) GetAllCompanies(ctx context.Context) ([]*CompanyResponse, error) {
	slog.Info("Getting all companies")

	companies, err := s.repo.GetAllCompanies(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*CompanyResponse, len(companies))
	for i, company := range companies {
		responses[i] = toCompanyResponse(company)
	}

	return responses, nil
}

func toCompanyResponse(company *Company) *CompanyResponse {
	return &CompanyResponse{
		ID:          company.ID,
		Name:        company.Name,
		Description: company.Description,
		Industry:    company.Industry,
		URL:         company.URL,
		Address:     company.Address,
		City:        company.City,
		State:       company.State,
		Country:     company.Country,
		ZipCode:     company.ZipCode,
		CreatedAt:   company.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   company.UpdatedAt.Format(time.RFC3339),
	}
}
