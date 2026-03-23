package repositories

import (
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

type CompanyRepository interface {
	Create(company entities.Company) (*entities.Company, error)
	CreateAddress(address entities.Address) (*entities.Address, error)
	FindAddressByCompanyID(companyId string) (*entities.Address, error)
	RegisterCompanyOwner(owner entities.CompanyOwner) error
	FindByCnpj(cnpj valueobjects.Cnpj) (*entities.Company, error)
	FindCompanyOwnerByEmail(email sharedvalueobjects.Email) (*entities.CompanyOwner, error)
	UpdateCompanyOwnerPassword(userID string, hashedPassword string) error
	InTransaction(fn func(CompanyRepository) error) error
}
