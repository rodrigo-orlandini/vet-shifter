package usecases

import (
	"rodrigoorlandini/vet-shifter/internal/shifts/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
)

type CreateShiftUseCase struct {
	repo repositories.ShiftRepository
}

type CreateShiftInput struct {
	Shift entities.Shift
}

type CreateShiftOutput struct {
	Shift *entities.Shift
}

func NewCreateShiftUseCase(repo repositories.ShiftRepository) *CreateShiftUseCase {
	return &CreateShiftUseCase{repo: repo}
}

func (u *CreateShiftUseCase) Execute(input *CreateShiftInput) (*CreateShiftOutput, error) {
	created, err := u.repo.Create(input.Shift)
	if err != nil {
		return nil, err
	}
	return &CreateShiftOutput{Shift: created}, nil
}
