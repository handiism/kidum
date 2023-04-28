package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	Form *formRepository
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Form: NewFormRepository(pool),
	}
}
