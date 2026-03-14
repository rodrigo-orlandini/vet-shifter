package mappers

import (
	"time"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

func CompanyOwnerFromPersistence(owner queries.CompanyOwner) (*entities.CompanyOwner, error) {
	email, err := sharedvalueobjects.NewEmail(owner.Email)
	if err != nil {
		return nil, err
	}

	phone, err := sharedvalueobjects.NewPhone(owner.Phone)
	if err != nil {
		return nil, err
	}

	o := &entities.CompanyOwner{
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

type CompanyOwnerFromHttpInput struct {
	Email         string
	Phone         string
	Password      string
	CompanyId     string
	ConsentLgpdAt *time.Time
}

func CompanyOwnerFromHttp(input CompanyOwnerFromHttpInput) (*entities.CompanyOwner, error) {
	email, err := sharedvalueobjects.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	phone, err := sharedvalueobjects.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	return entities.NewCompanyOwner(*email, *phone, input.Password, input.CompanyId, input.ConsentLgpdAt)
}
