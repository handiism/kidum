package domain

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/slices"
)

type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

func NewUserGender(str string) (Gender, error) {
	genders := []Gender{Male, Female}
	found := slices.Contains(genders, Gender(str))
	if !found {
		return "", errors.New("unknown form type")
	}

	return Gender(str), nil
}

type User struct {
	Id      int64
	Name    string
	Age     int64
	Address string
	Gender  Gender

	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt pgtype.Timestamptz

	FormId       sql.NullInt64
	CredentialId sql.NullInt64
}

type UserRequest struct {
	Name       string
	Age        int64
	Address    string
	Gender     Gender
	Credential CredentialRequest
}
