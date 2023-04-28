package domain

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/slices"
)

type FormStatus string

const (
	FormStatusIsUnproccessed FormStatus = "UNPROCCESSED"
	FormStatusIsDataInvalid  FormStatus = "DATA_INVALID"
	FormStatusIsDataValid    FormStatus = "DATA_VALID"
	FormStatusIsDone         FormStatus = "DONE"
)

type FormType string

const (
	FormTypeIndividual FormType = "INDIVIDUAL"
	FormTypeGroup      FormType = "GROUP"
)

func NewFormType(str string) (FormType, error) {
	types := []FormType{FormTypeIndividual, FormTypeGroup}
	found := slices.Contains(types, FormType(str))
	if !found {
		return "", errors.New("unknown form type")
	}

	return FormType(str), nil
}

type Form struct {
	Id          int64
	Code        string
	Agreement   bool
	Ticket      sql.NullString
	Description sql.NullString
	Status      FormStatus
	Type        FormType

	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt pgtype.Timestamptz

	DestinationRouteId sql.NullInt64
	AdminScannerId     sql.NullInt64
	AdminValidatorId   sql.NullInt64
	ContactId          sql.NullInt64
}

type FormRequest struct {
	Agreement          bool
	Type               FormType
	Users              []UserRequest
	DestinationRouteId int64
	Contact            ContactRequest
}

type FormRepository interface{}

type FormUsecase interface {
	Create(c context.Context, formRequest FormRequest)
}
