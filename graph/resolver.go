package graph

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Resolver struct {
	db *pgxpool.Pool
}

func NewResolver(db *pgxpool.Pool) *Resolver {
	return &Resolver{db: db}
}
