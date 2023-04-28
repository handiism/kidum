package usecase

import "github.com/handiism/kidum/repository"

type Usecase struct {
	Form *FormUsecase
}

func NewUsecase(repository *repository.Repository) Usecase {
	return Usecase{
		Form: NewFormUsecase(repository.Form),
	}
}
