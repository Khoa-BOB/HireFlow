package company

import (
	"context"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) CreateCompany(ctx context.Context, company *Company) (*Company, error) {
	query := `
		INSERT INTO companies (name, description, industry, url, address, city, state, country, zip_code)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`
	var id string
	var createdAt, updatedAt time.Time
	err := r.db.QueryRow(ctx, query,
		company.Name,
		company.Description,
		company.Industry,
		company.URL,
		company.Address,
		company.City,
		company.State,
		company.Country,
		company.ZipCode,
	).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return nil, err
	}

	return &Company{
		ID:          id,
		Name:        company.Name,
		Description: company.Description,
		Industry:    company.Industry,
		URL:         company.URL,
		Address:     company.Address,
		City:        company.City,
		State:       company.State,
		Country:     company.Country,
		ZipCode:     company.ZipCode,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func (r *PostgresRepository) UpdateCompany(ctx context.Context, id string, req *CompanyUpdate) (*Company, error) {
	query := `
		UPDATE companies
		SET
			name = COALESCE($2, name),
			description = COALESCE($3, description),
			url = COALESCE($4, url),
			industry = COALESCE($5, industry),
			address = COALESCE($6, address),
			city = COALESCE($7, city),
			state = COALESCE($8, state),
			country = COALESCE($9, country),
			zip_code = COALESCE($10, zip_code),
			updated_at = now()
		WHERE id = $1
		RETURNING
			id,
			name,
			description,
			url,
			industry,
			address,
			city,
			state,
			country,
			zip_code,
			created_at,
			updated_at;
	`
	
	var company Company

	err := r.db.QueryRow(
		ctx,
		query,
		id,
		req.Name,
		req.Description,
		req.URL,
		req.Industry,
		req.Address,
		req.City,
		req.State,
		req.Country,
		req.ZipCode,
	).Scan(
		&company.ID,
		&company.Name,
		&company.Description,
		&company.URL,
		&company.Industry,
		&company.Address,
		&company.City,
		&company.State,
		&company.Country,
		&company.ZipCode,
		&company.CreatedAt,
		&company.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (r *PostgresRepository) DeleteCompany(ctx context.Context, id string) error {
	query := `
		DELETE FROM companies
		WHERE id = $1;
	`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetCompanyByID(ctx context.Context, id string) (*Company, error) {
	query := `
		SELECT id, name, description, url, industry, address, city, state, country, zip_code, created_at, updated_at
		FROM companies
		WHERE id = $1;
	`

	var company Company

	err := r.db.QueryRow(ctx, query, id).Scan(
		&company.ID,
		&company.Name,
		&company.Description,
		&company.URL,
		&company.Industry,
		&company.Address,
		&company.City,
		&company.State,
		&company.Country,
		&company.ZipCode,
		&company.CreatedAt,
		&company.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (r *PostgresRepository) GetAllCompanies(ctx context.Context) ([]*Company, error) {
	query := `
		SELECT id, name, description, url, industry, address, city, state, country, zip_code, created_at, updated_at
		FROM companies;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []*Company

	for rows.Next() {
		var company Company
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Description,
			&company.URL,
			&company.Industry,
			&company.Address,
			&company.City,
			&company.State,
			&company.Country,
			&company.ZipCode,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		companies = append(companies, &company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

