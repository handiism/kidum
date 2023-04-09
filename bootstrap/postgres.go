package bootstrap

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgPool(env *Env) *pgxpool.Pool {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, env.PostgresURL)
	if err != nil {
		log.Fatal("Can't create postgres pool: ", err)
	}

	if _, err := pool.Acquire(ctx); err != nil {
		log.Fatal("Can't acquire postgres connection: ", err)
	}

	return pool
}
