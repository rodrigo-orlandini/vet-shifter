package mappers

import (
	"database/sql"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
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
		Street:         nullStringToStr(company.Street),
		Number:         nullStringToStr(company.Number),
		City:           nullStringToStr(company.City),
		State:          nullStringToStr(company.State),
		ZipCode:        nullStringToStr(company.ZipCode),
		ApprovalStatus: string(company.ApprovalStatus),
		CreatedAt:      &company.CreatedAt,
	}, nil
}

type CompanyFromHttpInput struct {
	Cnpj        string
	CompanyName string
	Street      string
	Number      string
	City        string
	State       string
	ZipCode     string
}

func CompanyFromHttp(input CompanyFromHttpInput) (*entities.Company, error) {
	cnpj, err := valueobjects.NewCnpj(input.Cnpj)
	if err != nil {
		return nil, err
	}
	var addr *entities.Address
	if input.Street != "" || input.City != "" || input.ZipCode != "" {
		addr = &entities.Address{
			Street:  input.Street,
			Number:  input.Number,
			City:    input.City,
			State:   input.State,
			ZipCode: input.ZipCode,
		}
	}
	return entities.NewCompany(*cnpj, input.CompanyName, addr)
}
