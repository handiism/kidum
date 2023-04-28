package domain

import (
	"errors"
	"mime/multipart"

	"golang.org/x/exp/slices"
)

type CredentialType string

const (
	NationalId    CredentialType = "NATIONAL_ID"
	DriverLicense CredentialType = "DRIVER_LICENSE"
)

func NewCredentialsType(str string) (CredentialType, error) {
	types := []CredentialType{NationalId, DriverLicense}
	found := slices.Contains(types, CredentialType(str))
	if !found {
		return "", errors.New("unknown form type")
	}

	return CredentialType(str), nil
}

type Credential struct {
	Id     int64
	Image  string
	Number string
	Type   CredentialType
}

type CredentialRequest struct {
	Image  *multipart.FileHeader
	Number string
	Type   CredentialType
}
