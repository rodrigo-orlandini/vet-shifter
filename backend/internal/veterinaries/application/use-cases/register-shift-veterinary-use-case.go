package usecases

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
)

type RegisterShiftVeterinaryUseCase struct {
	shiftVeterinaryRepository repositories.ShiftVeterinaryRepository
}

type RegisterShiftVeterinaryUseCaseInput struct {
	Veterinary entities.ShiftVeterinary
}

type RegisterShiftVeterinaryUseCaseOutput struct {
	VeterinaryId string
}

func NewRegisterShiftVeterinaryUseCase(
	shiftVeterinaryRepository repositories.ShiftVeterinaryRepository,
) *RegisterShiftVeterinaryUseCase {
	return &RegisterShiftVeterinaryUseCase{
		shiftVeterinaryRepository: shiftVeterinaryRepository,
	}
}

func (u *RegisterShiftVeterinaryUseCase) Execute(
	input *RegisterShiftVeterinaryUseCaseInput,
) (*RegisterShiftVeterinaryUseCaseOutput, error) {
	if input.Veterinary.ConsentLgpdAt == nil {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Consentimento LGPD",
			Value: "",
		}
	}

	existingByCpf, err := u.shiftVeterinaryRepository.FindByCpf(input.Veterinary.Cpf)
	if err != nil {
		return nil, &customerror.RepositoryError{
			Entity: "Veterinário",
			Field:  "CPF",
			Err:    err,
		}
	}

	if existingByCpf != nil {
		return nil, &customerror.AlreadyExistsError{
			Field: "CPF",
			Value: input.Veterinary.Cpf.GetMasked(),
		}
	}

	existingByEmail, err := u.shiftVeterinaryRepository.FindByEmail(input.Veterinary.Email)
	if err != nil {
		return nil, &customerror.RepositoryError{
			Entity: "Veterinário",
			Field:  "Email",
			Err:    err,
		}
	}

	if existingByEmail != nil {
		return nil, &customerror.AlreadyExistsError{
			Field: "E-mail",
			Value: input.Veterinary.Email.GetValue(),
		}
	}

	existingByPhone, err := u.shiftVeterinaryRepository.FindByPhone(input.Veterinary.Phone)
	if err != nil {
		return nil, &customerror.RepositoryError{
			Entity: "Veterinário",
			Field:  "Telefone",
			Err:    err,
		}
	}

	if existingByPhone != nil {
		return nil, &customerror.AlreadyExistsError{
			Field: "Telefone",
			Value: input.Veterinary.Phone.GetValue(),
		}
	}

	passwordHash := utils.Argon2Hash(input.Veterinary.Password)
	input.Veterinary.Password = passwordHash

	created, err := u.shiftVeterinaryRepository.Create(input.Veterinary)
	if err != nil {
		return nil, &customerror.RepositoryError{
			Entity: "Veterinário",
			Field:  "Criação",
			Err:    err,
		}
	}

	return &RegisterShiftVeterinaryUseCaseOutput{VeterinaryId: created.Id}, nil
}
