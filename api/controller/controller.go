package controller

import (
	"github.com/handiism/kidum/repository"
	"github.com/handiism/kidum/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct {
	Form FormController
}

func NewController(pool *pgxpool.Pool) Controller {
	repo := repository.NewRepository(pool)
	ucase := usecase.NewUsecase(repo)

	return Controller{
		Form: FormController{
			formUsecase: ucase.Form,
		},
	}
}
