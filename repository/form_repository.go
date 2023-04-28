package repository

import "github.com/jackc/pgx/v5/pgxpool"

type formRepository struct {
	pool *pgxpool.Pool
}

func NewFormRepository(pool *pgxpool.Pool) *formRepository {
	return &formRepository{
		pool,
	}
}
