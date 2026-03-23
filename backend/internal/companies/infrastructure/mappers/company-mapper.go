package mappers

import (
	"database/sql"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
)

func nullStringToStr(n sql.NullString) string {
	if n.Valid {
		return n.String
	}
	return ""
}

func StrToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

func CompanyFromPersistence(company queries.Company) (*entities.Company, error) {
	cnpj, err := valueobjects.NewCnpj(company.Cnpj)
	if err != nil {
		return nil, err
	}

	return &entities.Company{
		Id:             company.ID.String(),
		Cnpj:           *cnpj,
		Name:           company.Name,
		ApprovalStatus: string(company.ApprovalStatus),
		CreatedAt:      &company.CreatedAt,
	}, nil
}

func CompanyFromHttp(cnpj string, companyName string) (*entities.Company, error) {
	cnpjVO, err := valueobjects.NewCnpj(cnpj)
	if err != nil {
		return nil, err
	}

	return entities.NewCompany(*cnpjVO, companyName)
}

func AddressFromPersistence(addr queries.Address) (*entities.Address, error) {
	uf, err := sharedvalueobjects.NewUF(nullStringToStr(addr.State))
	if err != nil {
		return nil, err
	}

	cep, err := sharedvalueobjects.NewCep(nullStringToStr(addr.ZipCode))
	if err != nil {
		return nil, err
	}

	return &entities.Address{
		Id:        addr.ID.String(),
		CompanyId: addr.CompanyID.String(),
		Street:    nullStringToStr(addr.Street),
		Number:    nullStringToStr(addr.Number),
		City:      nullStringToStr(addr.City),
		State:     *uf,
		ZipCode:   *cep,
		CreatedAt: &addr.CreatedAt,
	}, nil
}
