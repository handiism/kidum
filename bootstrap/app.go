package bootstrap

import "github.com/jackc/pgx/v5/pgxpool"

type Application struct {
	Env    *Env
	PgPool *pgxpool.Pool
}

func App() *Application {
	app := &Application{}

	app.Env = NewEnv()
	app.PgPool = NewPgPool(app.Env)

	return app
}

func (a *Application) Close() {
	a.PgPool.Close()
}
