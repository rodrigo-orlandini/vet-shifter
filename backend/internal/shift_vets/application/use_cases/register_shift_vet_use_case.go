package usecases

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/domain/entities"
)

type RegisterShiftVetUseCase struct {
	repo repositories.ShiftVetRepository
}

type RegisterShiftVetInput struct {
	Vet entities.ShiftVet
}

type RegisterShiftVetOutput struct {
	Vet *entities.ShiftVet
}

func NewRegisterShiftVetUseCase(repo repositories.ShiftVetRepository) *RegisterShiftVetUseCase {
	return &RegisterShiftVetUseCase{repo: repo}
}

func (u *RegisterShiftVetUseCase) Execute(input *RegisterShiftVetInput) (*RegisterShiftVetOutput, error) {
	existingEmail, err := u.repo.FindByEmail(input.Vet.Email)
	if err != nil {
		return nil, &customerror.RepositoryError{Entity: "ShiftVet", Field: "Email", Err: err}
	}
	if existingEmail != nil {
		return nil, &customerror.AlreadyExistsError{Entity: "ShiftVet", Field: "Email", Value: input.Vet.Email.GetValue()}
	}
	existingCpf, err := u.repo.FindByCpf(input.Vet.Cpf)
	if err != nil {
		return nil, &customerror.RepositoryError{Entity: "ShiftVet", Field: "Cpf", Err: err}
	}
	if existingCpf != nil {
		return nil, &customerror.AlreadyExistsError{Entity: "ShiftVet", Field: "Cpf", Value: input.Vet.Cpf.GetMasked()}
	}
	vet := input.Vet
	vet.Password = utils.Argon2Hash(vet.Password)
	created, err := u.repo.Create(vet)
	if err != nil {
		return nil, err
	}
	return &RegisterShiftVetOutput{Vet: created}, nil
}
