package entities

import (
	"time"

	"github.com/google/uuid"

	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

type CompanyOwner struct {
	Id            string
	Email         sharedvalueobjects.Email
	Phone         sharedvalueobjects.Phone
	Password      string
	CompanyId     string
	ConsentLgpdAt *time.Time
}

func NewCompanyOwner(email sharedvalueobjects.Email, phone sharedvalueobjects.Phone, password string, companyId string, consentLgpdAt *time.Time) (*CompanyOwner, error) {
	id, _ := uuid.NewV7()

	owner := &CompanyOwner{
		Id:            id.String(),
		Email:         email,
		Phone:         phone,
		Password:      password,
		CompanyId:     companyId,
		ConsentLgpdAt: consentLgpdAt,
	}

	return owner, nil
}
