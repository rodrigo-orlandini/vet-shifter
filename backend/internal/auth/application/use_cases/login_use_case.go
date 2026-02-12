package usecases

import (
	"time"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/internal/companies/application/repositories"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	shiftvetrepo "rodrigoorlandini/vet-shifter/internal/shift_vets/application/repositories"
)

const (
	RoleClinic = "clinic"
	RoleVet    = "vet"
)

const TokenExpiry = 24 * time.Hour

type LoginUseCase struct {
	companyRepo   repositories.CompanyRepository
	shiftVetRepo  shiftvetrepo.ShiftVetRepository
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token string
	Role  string
	Sub   string // owner_id or vet_id
}

func NewLoginUseCase(companyRepo repositories.CompanyRepository, shiftVetRepo shiftvetrepo.ShiftVetRepository) *LoginUseCase {
	return &LoginUseCase{companyRepo: companyRepo, shiftVetRepo: shiftVetRepo}
}

func (u *LoginUseCase) Execute(input *LoginInput) (*LoginOutput, error) {
	email, err := valueobjects.NewEmail(input.Email)
	if err != nil {
		return nil, &customerror.InvalidCredentialsError{}
	}
	owner, err := u.companyRepo.FindOwnerByEmail(*email)
	if err != nil {
		return nil, err
	}
	if owner != nil && utils.Argon2Compare(input.Password, owner.Password) {
		token, err := utils.JWTGenerate(owner.Id, RoleClinic, TokenExpiry)
		if err != nil {
			return nil, err
		}
		return &LoginOutput{Token: token, Role: RoleClinic, Sub: owner.Id}, nil
	}
	vet, err := u.shiftVetRepo.FindByEmail(*email)
	if err != nil {
		return nil, err
	}
	if vet != nil && utils.Argon2Compare(input.Password, vet.Password) {
		token, err := utils.JWTGenerate(vet.Id, RoleVet, TokenExpiry)
		if err != nil {
			return nil, err
		}
		return &LoginOutput{Token: token, Role: RoleVet, Sub: vet.Id}, nil
	}
	return nil, &customerror.InvalidCredentialsError{}
}
