package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvo "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/shift_vets/application/use_cases"
	shiftvetentities "rodrigoorlandini/vet-shifter/internal/shift_vets/domain/entities"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func TestRegisterShiftVetUseCase(t *testing.T) {
	t.Run("registers a new shift vet successfully", func(t *testing.T) {
		repo := repositories.NewStubShiftVetRepository()
		uc := usecases.NewRegisterShiftVetUseCase(repo)

		email, _ := valueobjects.NewEmail("vet@test.com")
		phone, _ := valueobjects.NewPhone("11999991111")
		cpf, _ := sharedvo.NewCpf("12345678901")
		vet, _ := shiftvetentities.NewShiftVet(*email, *phone, "plainpass", "Dr. João", *cpf, "12345", "SP", []string{"Clínico"}, nil)

		out, err := uc.Execute(&usecases.RegisterShiftVetInput{Vet: *vet})
		assert.Nil(t, err)
		assert.NotNil(t, out.Vet)
		assert.NotEmpty(t, out.Vet.Id)
		assert.NotEqual(t, "plainpass", out.Vet.Password)

		found, _ := repo.FindByEmail(*email)
		assert.NotNil(t, found)
		assert.Equal(t, "vet@test.com", found.Email.GetValue())
		assert.Equal(t, "12345678901", found.Cpf.GetValue())
	})

	t.Run("returns AlreadyExistsError when email already registered", func(t *testing.T) {
		repo := repositories.NewStubShiftVetRepository()
		uc := usecases.NewRegisterShiftVetUseCase(repo)

		email, _ := valueobjects.NewEmail("existing@test.com")
		phone, _ := valueobjects.NewPhone("11999992222")
		cpf, _ := sharedvo.NewCpf("12345678901")
		vet, _ := shiftvetentities.NewShiftVet(*email, *phone, "pass", "Dr. Exist", *cpf, "11111", "SP", nil, nil)
		repo.Create(*vet)

		cpf2, _ := sharedvo.NewCpf("98765432100")
		vet2, _ := shiftvetentities.NewShiftVet(*email, *phone, "other", "Dr. New", *cpf2, "22222", "RJ", nil, nil)

		_, err := uc.Execute(&usecases.RegisterShiftVetInput{Vet: *vet2})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.AlreadyExistsError{}, err)
		assert.Equal(t, "ShiftVet", err.(*customerror.AlreadyExistsError).Entity)
		assert.Equal(t, "Email", err.(*customerror.AlreadyExistsError).Field)
	})

	t.Run("returns AlreadyExistsError when CPF already registered", func(t *testing.T) {
		repo := repositories.NewStubShiftVetRepository()
		uc := usecases.NewRegisterShiftVetUseCase(repo)

		email1, _ := valueobjects.NewEmail("first@test.com")
		phone, _ := valueobjects.NewPhone("11999993333")
		cpf, _ := sharedvo.NewCpf("11111111111")
		vet1, _ := shiftvetentities.NewShiftVet(*email1, *phone, "pass", "Dr. First", *cpf, "33333", "SP", nil, nil)
		repo.Create(*vet1)

		email2, _ := valueobjects.NewEmail("second@test.com")
		vet2, _ := shiftvetentities.NewShiftVet(*email2, *phone, "pass2", "Dr. Second", *cpf, "44444", "MG", nil, nil)

		_, err := uc.Execute(&usecases.RegisterShiftVetInput{Vet: *vet2})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.AlreadyExistsError{}, err)
		assert.Equal(t, "Cpf", err.(*customerror.AlreadyExistsError).Field)
	})
}
