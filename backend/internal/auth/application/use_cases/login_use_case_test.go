package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvo "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use_cases"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	shiftvetentities "rodrigoorlandini/vet-shifter/internal/shift_vets/domain/entities"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func TestLoginUseCase(t *testing.T) {
	t.Run("returns token and role clinic when owner exists and password matches", func(t *testing.T) {
		companyRepo := repositories.NewStubCompanyRepository()
		shiftVetRepo := repositories.NewStubShiftVetRepository()
		uc := usecases.NewLoginUseCase(companyRepo, shiftVetRepo)

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		company, _ := entities.NewCompany(*cnpj, "Clinic", nil)
		companyRepo.Create(*company)

		email, _ := valueobjects.NewEmail("owner@clinic.com")
		phone, _ := valueobjects.NewPhone("11999990000")
		owner, _ := entities.NewOwner(*email, *phone, utils.Argon2Hash("secret"), company.Id, nil)
		companyRepo.RegisterOwner(*owner)

		out, err := uc.Execute(&usecases.LoginInput{Email: "owner@clinic.com", Password: "secret"})
		assert.Nil(t, err)
		assert.NotEmpty(t, out.Token)
		assert.Equal(t, usecases.RoleClinic, out.Role)
		assert.Equal(t, owner.Id, out.Sub)
	})

	t.Run("returns token and role vet when shift vet exists and password matches", func(t *testing.T) {
		companyRepo := repositories.NewStubCompanyRepository()
		shiftVetRepo := repositories.NewStubShiftVetRepository()
		uc := usecases.NewLoginUseCase(companyRepo, shiftVetRepo)

		email, _ := valueobjects.NewEmail("vet@email.com")
		phone, _ := valueobjects.NewPhone("11988880000")
		cpf, _ := sharedvo.NewCpf("12345678901")
		vet, _ := shiftvetentities.NewShiftVet(*email, *phone, utils.Argon2Hash("vetpass"), "Dr. Vet", *cpf, "12345", "SP", nil, nil)
		shiftVetRepo.Create(*vet)

		out, err := uc.Execute(&usecases.LoginInput{Email: "vet@email.com", Password: "vetpass"})
		assert.Nil(t, err)
		assert.NotEmpty(t, out.Token)
		assert.Equal(t, usecases.RoleVet, out.Role)
		assert.Equal(t, vet.Id, out.Sub)
	})

	t.Run("returns InvalidCredentialsError when password is wrong", func(t *testing.T) {
		companyRepo := repositories.NewStubCompanyRepository()
		shiftVetRepo := repositories.NewStubShiftVetRepository()
		uc := usecases.NewLoginUseCase(companyRepo, shiftVetRepo)

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		company, _ := entities.NewCompany(*cnpj, "Clinic", nil)
		companyRepo.Create(*company)
		email, _ := valueobjects.NewEmail("owner@clinic.com")
		phone, _ := valueobjects.NewPhone("11999990000")
		owner, _ := entities.NewOwner(*email, *phone, utils.Argon2Hash("secret"), company.Id, nil)
		companyRepo.RegisterOwner(*owner)

		_, err := uc.Execute(&usecases.LoginInput{Email: "owner@clinic.com", Password: "wrong"})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})

	t.Run("returns InvalidCredentialsError when email does not exist", func(t *testing.T) {
		companyRepo := repositories.NewStubCompanyRepository()
		shiftVetRepo := repositories.NewStubShiftVetRepository()
		uc := usecases.NewLoginUseCase(companyRepo, shiftVetRepo)

		_, err := uc.Execute(&usecases.LoginInput{Email: "nobody@example.com", Password: "any"})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})
}
