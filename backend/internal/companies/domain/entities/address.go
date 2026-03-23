package entities

import (
	"time"

	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"

	"github.com/google/uuid"
)

type Address struct {
	Id        string
	CompanyId string
	Street    string
	Number    string
	City      string
	State     sharedvalueobjects.UF
	ZipCode   sharedvalueobjects.Cep
	CreatedAt *time.Time
}

func NewAddress(companyId, street, number, city, state, zipCode string) (*Address, error) {
	uf, err := sharedvalueobjects.NewUF(state)
	if err != nil {
		return nil, err
	}

	cep, err := sharedvalueobjects.NewCep(zipCode)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	id, _ := uuid.NewV7()

	return &Address{
		Id:        id.String(),
		CompanyId: companyId,
		Street:    street,
		Number:    number,
		City:      city,
		State:     *uf,
		ZipCode:   *cep,
		CreatedAt: &now,
	}, nil
}
