package usecase

import "github.com/handiism/kidum/domain"

type FormUsecase struct {
	FormRepository domain.FormRepository
}

func NewFormUsecase(formRepository domain.FormRepository) *FormUsecase {
	return &FormUsecase{
		formRepository,
	}
}
