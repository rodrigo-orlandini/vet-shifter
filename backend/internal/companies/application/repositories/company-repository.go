package repositories

import (
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
)

type CompanyRepository interface {
	Create(company entities.Company) (*entities.Company, error)
	RegisterOwner(owner entities.Owner) error
	FindByCnpj(cnpj valueobjects.Cnpj) (*entities.Company, error)
	FindOwnerByEmail(email valueobjects.Email) (*entities.Owner, error)
}
