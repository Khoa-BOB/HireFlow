package auth

import (
	"context"
	"errors"
	"time"

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

func (r *PostgresRepository) GetUsers(ctx context.Context) ([]UserWithRole, error){
	query := `
			SELECT
				u.id,
				u.email,
				u.full_name,
				r.name AS role,
				u.created_at
			FROM users u
			JOIN roles r
				ON u.role_id = r.id
			ORDER BY u.created_at DESC;
			`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []UserWithRole{}

	for rows.Next(){
		var user UserWithRole

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FullName,
			&user.Role,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users,nil

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

func (r *PostgresRepository) CreateRefreshToken(ctx context.Context, id, refreshToken string, expiresAt time.Time) (*RefreshToken, error){
	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, token, created_at, expires_at, revoked_at;
	`
	var token RefreshToken

	err := r.db.QueryRow(ctx, query, id, refreshToken, expiresAt).Scan(
		&token.ID,
		&token.UserID,
		&token.Token,
		&token.CreatedAt,
		&token.ExpiresAt,
		&token.RevokedAt,
	)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *PostgresRepository) GetRefreshTokenByHash(ctx context.Context, refreshTokenHash string) (*RefreshToken, error){
	query := `
		SELECT id, user_id, token, created_at, expires_at, revoked_at
		FROM refresh_tokens
		WHERE token = $1
	`
	var token RefreshToken

	err := r.db.QueryRow(ctx, query, refreshTokenHash).Scan(
		&token.ID,
		&token.UserID,
		&token.Token,
		&token.CreatedAt,
		&token.ExpiresAt,
		&token.RevokedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *PostgresRepository) RevokeRefreshToken(ctx context.Context, refreshTokenHash string) error{
	query := `
		UPDATE refresh_tokens
		SET revoked_at = NOW()
		WHERE token = $1 AND revoked_at IS NULL;
	`

	cmdTag, err := r.db.Exec(ctx, query, refreshTokenHash)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("refresh token not found or already revoked")
	}

	return nil
}