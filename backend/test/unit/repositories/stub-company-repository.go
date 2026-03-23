package repositories

import (
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	companyrepositories "rodrigoorlandini/vet-shifter/internal/companies/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
)

type StubCompanyRepository struct {
	companies     map[string]*entities.Company
	companyOwners map[string]*entities.CompanyOwner
	addresses     map[string]*entities.Address
}

func NewStubCompanyRepository() *StubCompanyRepository {
	return &StubCompanyRepository{
		companies:     make(map[string]*entities.Company),
		companyOwners: make(map[string]*entities.CompanyOwner),
		addresses:     make(map[string]*entities.Address),
	}
}

func (r *StubCompanyRepository) Create(company entities.Company) (*entities.Company, error) {
	r.companies[company.Id] = &company

	return &company, nil
}

func (r *StubCompanyRepository) CreateAddress(address entities.Address) (*entities.Address, error) {
	r.addresses[address.CompanyId] = &address

	return &address, nil
}

func (r *StubCompanyRepository) FindAddressByCompanyID(companyId string) (*entities.Address, error) {
	if addr, ok := r.addresses[companyId]; ok {
		return addr, nil
	}

	return nil, nil
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

func (r *StubCompanyRepository) InTransaction(fn func(companyrepositories.CompanyRepository) error) error {
	return fn(r)
}
