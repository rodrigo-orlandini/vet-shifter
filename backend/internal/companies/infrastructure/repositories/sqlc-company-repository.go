package repositories

import (
	"context"
	"database/sql"
	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"rodrigoorlandini/vet-shifter/internal/companies/infrastructure/mappers"

	"github.com/google/uuid"
)

type SqlcCompanyRepository struct {
	queries *queries.Queries
}

func NewSqlcCompanyRepository() *SqlcCompanyRepository {
	conn := database.GetPostgresConnection()
	q := database.NewQueries(conn)

	return &SqlcCompanyRepository{
		queries: q,
	}
}

func (r *SqlcCompanyRepository) Create(company entities.Company) (*entities.Company, error) {
	ctx := context.Background()
	params := queries.CreateCompanyParams{
		ID:             uuid.MustParse(company.Id),
		Cnpj:           company.Cnpj.GetValue(),
		Name:           company.Name,
		Street:         mappers.StrToNullString(company.Street),
		Number:         mappers.StrToNullString(company.Number),
		City:           mappers.StrToNullString(company.City),
		State:          mappers.StrToNullString(company.State),
		ZipCode:        mappers.StrToNullString(company.ZipCode),
		ApprovalStatus: company.ApprovalStatus,
	}
	createdCompany, err := r.queries.CreateCompany(ctx, params)
	if err != nil {
		return nil, err
	}

	mappedCompany, err := mappers.CompanyFromPersistence(createdCompany)
	if err != nil {
		return nil, err
	}

	return mappedCompany, nil
}

func (r *SqlcCompanyRepository) RegisterOwner(owner entities.Owner) error {
	ctx := context.Background()
	consentLgpdAt := sql.NullTime{}
	if owner.ConsentLgpdAt != nil {
		consentLgpdAt = sql.NullTime{Time: *owner.ConsentLgpdAt, Valid: true}
	}
	err := r.queries.RegisterCompanyOwner(ctx, queries.RegisterCompanyOwnerParams{
		ID:            uuid.MustParse(owner.Id),
		Email:         owner.Email.GetValue(),
		Phone:         owner.Phone.GetValue(),
		Password:      owner.Password,
		CompanyID:     uuid.MustParse(owner.CompanyId),
		ConsentLgpdAt: consentLgpdAt,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlcCompanyRepository) FindByCnpj(cnpj valueobjects.Cnpj) (*entities.Company, error) {
	ctx := context.Background()

	foundCompany, err := r.queries.FindCompanyByCnpj(ctx, cnpj.GetValue())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	mappedCompany, err := mappers.CompanyFromPersistence(foundCompany)
	if err != nil {
		return nil, err
	}

	return mappedCompany, nil
}

func (r *SqlcCompanyRepository) FindOwnerByEmail(email valueobjects.Email) (*entities.Owner, error) {
	ctx := context.Background()

	foundOwner, err := r.queries.FindCompanyOwnerByEmail(ctx, email.GetValue())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	mappedOwner, err := mappers.OwnerFromPersistence(foundOwner)
	if err != nil {
		return nil, err
	}

	return mappedOwner, nil
}
