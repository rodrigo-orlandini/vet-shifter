package repositories

import (
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

type StubCompanyRepository struct {
	companies     map[string]*entities.Company
	companyOwners map[string]*entities.CompanyOwner
}

func NewStubCompanyRepository() *StubCompanyRepository {
	return &StubCompanyRepository{
		companies:     make(map[string]*entities.Company),
		companyOwners: make(map[string]*entities.CompanyOwner),
	}
}

func (r *StubCompanyRepository) Create(company entities.Company) (*entities.Company, error) {
	r.companies[company.Id] = &company

	return &company, nil
}

func (r *StubCompanyRepository) RegisterCompanyOwner(owner entities.CompanyOwner) error {
	r.companyOwners[owner.Id] = &owner

	return nil
}

func (r *StubCompanyRepository) FindByCnpj(cnpj valueobjects.Cnpj) (*entities.Company, error) {
	for _, company := range r.companies {
		if company.Cnpj.GetValue() == cnpj.GetValue() {
			return company, nil
		}
	}

	return nil, nil
}

func (r *StubCompanyRepository) FindCompanyOwnerByEmail(email sharedvalueobjects.Email) (*entities.CompanyOwner, error) {
	for _, owner := range r.companyOwners {
		if owner.Email.GetValue() == email.GetValue() {
			return owner, nil
		}
	}

	return nil, nil
}

func (r *StubCompanyRepository) UpdateCompanyOwnerPassword(userID string, hashedPassword string) error {
	for _, owner := range r.companyOwners {
		if owner.Id == userID {
			owner.Password = hashedPassword
			return nil
		}
	}

	return nil
}
