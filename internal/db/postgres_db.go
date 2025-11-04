package db

import (
	"context"
	"equi_genea_horse_service/config"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg *config.DatabaseConfig) (*Queries, *pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, nil, err
	}

	queries := New(pool)

	return queries, pool, nil
}
