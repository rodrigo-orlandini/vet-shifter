package entities

import (
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"time"

	"github.com/google/uuid"
)

type Owner struct {
	Id            string
	Email         valueobjects.Email
	Phone         valueobjects.Phone
	Password      string
	CompanyId     string
	ConsentLgpdAt *time.Time
}

func NewOwner(email valueobjects.Email, phone valueobjects.Phone, password string, companyId string, consentLgpdAt *time.Time) (*Owner, error) {
	id, _ := uuid.NewV7()
	owner := &Owner{
		Id:            id.String(),
		Email:         email,
		Phone:         phone,
		Password:      password,
		CompanyId:     companyId,
		ConsentLgpdAt: consentLgpdAt,
	}
	return owner, nil
}
