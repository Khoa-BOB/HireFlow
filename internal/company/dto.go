package company

import ()

type CreateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    string `json:"industry"`
	URL         string `json:"url"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	ZipCode     string `json:"zip_code"`
}

type UpdateCompanyRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Industry    *string `json:"industry"`
	URL         *string `json:"url"`
	Address     *string `json:"address"`
	City        *string `json:"city"`
	State       *string `json:"state"`
	Country     *string `json:"country"`
	ZipCode     *string `json:"zip_code"`
}

type CompanyResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    string `json:"industry"`
	URL         string `json:"url"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	ZipCode     string `json:"zip_code"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

