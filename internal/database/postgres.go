package database

import (
	"context"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresConnection() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}