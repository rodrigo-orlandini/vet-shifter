package repositories

import (
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
)

type StubCompanyRepository struct {
	companies map[string]*entities.Company
	owners    map[string]*entities.Owner
}

func NewStubCompanyRepository() *StubCompanyRepository {
	return &StubCompanyRepository{
		companies: make(map[string]*entities.Company),
		owners:    make(map[string]*entities.Owner),
	}
}

func (r *StubCompanyRepository) Create(company entities.Company) (*entities.Company, error) {
	r.companies[company.Id] = &company
	return &company, nil
}

func (r *StubCompanyRepository) RegisterOwner(owner entities.Owner) error {
	r.owners[owner.Id] = &owner
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

func (r *StubCompanyRepository) FindOwnerByEmail(email valueobjects.Email) (*entities.Owner, error) {
	for _, owner := range r.owners {
		if owner.Email.GetValue() == email.GetValue() {
			return owner, nil
		}
	}
	return nil, nil
}
