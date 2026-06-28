package auth

import (
	"context"
	"log/slog"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct{
	db *pgxpool.Pool
}

func NewPostgresRepository(db * pgxpool.Pool) * PostgresRepository{
	return &PostgresRepository{db:db}
}


func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*User, error){
	query := `
		SELECT id, email, full_name, password_hash, (SELECT name FROM roles WHERE id = role_id) AS role, created_at
		FROM users
		WHERE email = $1
	`
	var user User

	slog.Info("Query database", "query", query)
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.FullName,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user,nil
}

func (r *PostgresRepository)  GetUserByID(ctx context.Context, uuid string) (*User, error) {
	query := `
		SELECT id, email, full_name, password_hash, (SELECT name FROM roles WHERE id = role_id) AS role, created_at
		FROM users
		WHERE id = $1
	`
	var user User

	slog.Info("Query database", "query", query)
	err := r.db.QueryRow(ctx, query, uuid).Scan(
		&user.ID,
		&user.Email,
		&user.FullName,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user,nil
}

func (r *PostgresRepository) CreateUser(ctx context.Context,
										email string,
										passwordHash string,
										fullname string) (*User, error){
	query := `
		WITH inserted AS (
			INSERT INTO users (email, password_hash, full_name, role_id)
			VALUES ($1, $2, $3, (SELECT id FROM roles WHERE name = 'candidate'))
			RETURNING id, email, password_hash, full_name, role_id, created_at
		)
		SELECT inserted.id, inserted.email, inserted.password_hash, inserted.full_name, roles.name AS role, inserted.created_at
		FROM inserted
		JOIN roles ON roles.id = inserted.role_id;
	`
	var user User

	slog.Info("Query database", "query", query)
	err := r.db.QueryRow(ctx, query, email, passwordHash, fullname).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}