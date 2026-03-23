package repositories

import (
	"context"
	"database/sql"
	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	companyrepositories "rodrigoorlandini/vet-shifter/internal/companies/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"rodrigoorlandini/vet-shifter/internal/companies/infrastructure/mappers"

	"github.com/google/uuid"
)

type SqlcCompanyRepository struct {
	db      *sql.DB
	queries *queries.Queries
}

func NewSqlcCompanyRepository() *SqlcCompanyRepository {
	conn := database.GetPostgresConnection()
	q := database.NewQueries(conn)

	return &SqlcCompanyRepository{
		db:      conn,
		queries: q,
	}
}

func (r *SqlcCompanyRepository) Create(company entities.Company) (*entities.Company, error) {
	ctx := context.Background()
	params := queries.CreateCompanyParams{
		ID:             uuid.MustParse(company.Id),
		Cnpj:           company.Cnpj.GetValue(),
		Name:           company.Name,
		ApprovalStatus: queries.AccountStatus(company.ApprovalStatus),
	}

	createdCompany, err := r.queries.CreateCompany(ctx, params)
	if err != nil {
		return nil, err
	}

	return mappers.CompanyFromPersistence(createdCompany)
}

func (r *SqlcCompanyRepository) CreateAddress(address entities.Address) (*entities.Address, error) {
	ctx := context.Background()
	created, err := r.queries.CreateAddress(ctx, queries.CreateAddressParams{
		ID:        uuid.MustParse(address.Id),
		CompanyID: uuid.MustParse(address.CompanyId),
		Street:    mappers.StrToNullString(address.Street),
		Number:    mappers.StrToNullString(address.Number),
		City:      mappers.StrToNullString(address.City),
		State:     mappers.StrToNullString(address.State.GetValue()),
		ZipCode:   mappers.StrToNullString(address.ZipCode.GetValue()),
	})
	if err != nil {
		return nil, err
	}

	return mappers.AddressFromPersistence(created)
}

func (r *SqlcCompanyRepository) FindAddressByCompanyID(companyId string) (*entities.Address, error) {
	ctx := context.Background()
	addr, err := r.queries.FindAddressByCompanyID(ctx, uuid.MustParse(companyId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return mappers.AddressFromPersistence(addr)
}

func (r *SqlcCompanyRepository) RegisterCompanyOwner(owner entities.CompanyOwner) error {
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

func (r *SqlcCompanyRepository) FindCompanyOwnerByEmail(email sharedvalueobjects.Email) (*entities.CompanyOwner, error) {
	ctx := context.Background()
	foundOwner, err := r.queries.FindCompanyOwnerByEmail(ctx, email.GetValue())

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return mappers.CompanyOwnerFromPersistence(foundOwner)
}

func (r *SqlcCompanyRepository) UpdateCompanyOwnerPassword(userID string, hashedPassword string) error {
	ctx := context.Background()
	id := uuid.MustParse(userID)

	return r.queries.UpdateCompanyOwnerPassword(ctx, queries.UpdateCompanyOwnerPasswordParams{
		ID:       id,
		Password: hashedPassword,
	})
}

func (r *SqlcCompanyRepository) InTransaction(fn func(companyrepositories.CompanyRepository) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	txRepo := &SqlcCompanyRepository{
		db:      r.db,
		queries: r.queries.WithTx(tx),
	}

	if err := fn(txRepo); err != nil {
		return err
	}

	return tx.Commit()
}
