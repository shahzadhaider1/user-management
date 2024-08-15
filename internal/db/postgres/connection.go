package db

import (
	"context"
	"fmt"

	"user-management/internal/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	Queries *db.Queries
}

// NewStore initializes a new Store with the provided connection pool
func NewStore(pool db.DBTX) *Store {
	return &Store{
		Queries: db.New(pool),
	}
}

// OpenDB creates a new connection pool to the PostgreSQL database
func OpenDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil
}

// CloseDB closes the database connection pool
func CloseDB(pool *pgxpool.Pool) {
	pool.Close()
}
