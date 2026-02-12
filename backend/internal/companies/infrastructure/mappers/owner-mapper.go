package mappers

import (
	"time"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
)

func OwnerFromPersistence(owner queries.CompanyOwner) (*entities.Owner, error) {
	email, err := valueobjects.NewEmail(owner.Email)
	if err != nil {
		return nil, err
	}

	phone, err := valueobjects.NewPhone(owner.Phone)
	if err != nil {
		return nil, err
	}

	o := &entities.Owner{
		Id:        owner.ID.String(),
		Email:     *email,
		Phone:     *phone,
		Password:  owner.Password,
		CompanyId: owner.CompanyID.String(),
	}
	if owner.ConsentLgpdAt.Valid {
		o.ConsentLgpdAt = &owner.ConsentLgpdAt.Time
	}
	return o, nil
}

type OwnerFromHttpInput struct {
	Email         string
	Phone         string
	Password      string
	CompanyId     string
	ConsentLgpdAt *time.Time
}

func OwnerFromHttp(input OwnerFromHttpInput) (*entities.Owner, error) {
	email, err := valueobjects.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}
	phone, err := valueobjects.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}
	return entities.NewOwner(*email, *phone, input.Password, input.CompanyId, input.ConsentLgpdAt)
}
