package usecases_test

import (
	"testing"
	"time"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/veterinaries/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
	"rodrigoorlandini/vet-shifter/test/unit/factories"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseRegisterShiftVeterinary(t *testing.T) {
	t.Run("it should be able to register a new shift veterinary", func(t *testing.T) {
		useCase, deps := factories.NewRegisterShiftVeterinaryStubFactory()
		consent := time.Now()

		email, _ := sharedvalueobjects.NewEmail("vet@email.com")
		phone, _ := sharedvalueobjects.NewPhone("11999999999")
		cpf, _ := sharedvalueobjects.NewCpf("12345678901")
		crmv, _ := valueobjects.NewCrmv("12345", "SP")
		specialties, _ := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyGeneralPractice})

		v, err := entities.NewShiftVeterinary(
			*email,
			*phone,
			"password123",
			"Vet Name",
			*cpf,
			*crmv,
			*specialties,
			*valueobjects.PendingDocumentApproval(),
			&consent,
		)
		assert.Nil(t, err)

		out, err := useCase.Execute(&usecases.RegisterShiftVeterinaryUseCaseInput{
			Veterinary: *v,
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, out.VeterinaryId)

		found, err := deps.ShiftVeterinaryRepository.FindByEmail(*email)
		assert.Nil(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, "Vet Name", found.FullName)
		assert.Equal(t, "12345678901", found.Cpf.GetValue())
		assert.Equal(t, "vet@email.com", found.Email.GetValue())
		assert.NotEqual(t, "password123", found.Password)
	})

	t.Run("it should not be able to register when CPF already exists", func(t *testing.T) {
		useCase, _ := factories.NewRegisterShiftVeterinaryStubFactory()
		consent := time.Now()

		email1, _ := sharedvalueobjects.NewEmail("existing@email.com")
		phone1, _ := sharedvalueobjects.NewPhone("11999999999")
		cpf, _ := sharedvalueobjects.NewCpf("12345678901")
		crmv1, _ := valueobjects.NewCrmv("12345", "SP")
		specialties1, _ := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyGeneralPractice})
		existing, _ := entities.NewShiftVeterinary(
			*email1,
			*phone1,
			"password123",
			"Existing Vet",
			*cpf,
			*crmv1,
			*specialties1,
			*valueobjects.PendingDocumentApproval(),
			&consent,
		)

		_, err := useCase.Execute(&usecases.RegisterShiftVeterinaryUseCaseInput{Veterinary: *existing})
		assert.Nil(t, err)

		email2, _ := sharedvalueobjects.NewEmail("duplicate@email.com")
		phone2, _ := sharedvalueobjects.NewPhone("11888888888")
		crmv2, _ := valueobjects.NewCrmv("54321", "RJ")
		specialties2, _ := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyGeneralPractice})
		duplicate, _ := entities.NewShiftVeterinary(
			*email2,
			*phone2,
			"password123",
			"Duplicate Vet",
			*cpf,
			*crmv2,
			*specialties2,
			*valueobjects.PendingDocumentApproval(),
			&consent,
		)

		_, err = useCase.Execute(&usecases.RegisterShiftVeterinaryUseCaseInput{Veterinary: *duplicate})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.AlreadyExistsError{}, err)
		alreadyExistsErr := err.(*customerror.AlreadyExistsError)
		assert.Equal(t, "CPF", alreadyExistsErr.Field)
	})

	t.Run("it should not be able to register without LGPD consent", func(t *testing.T) {
		useCase, _ := factories.NewRegisterShiftVeterinaryStubFactory()
		consent := time.Now()

		email, _ := sharedvalueobjects.NewEmail("vet@email.com")
		phone, _ := sharedvalueobjects.NewPhone("11999999999")
		cpf, _ := sharedvalueobjects.NewCpf("12345678901")
		crmv, _ := valueobjects.NewCrmv("12345", "SP")
		specialties, _ := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyGeneralPractice})

		v, err := entities.NewShiftVeterinary(
			*email, *phone, "password123", "Vet Name",
			*cpf, *crmv, *specialties,
			*valueobjects.PendingDocumentApproval(), &consent,
		)
		assert.Nil(t, err)

		v.ConsentLgpdAt = nil

		_, err = useCase.Execute(&usecases.RegisterShiftVeterinaryUseCaseInput{Veterinary: *v})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidValueObjectError{}, err)
		invalidErr := err.(*customerror.InvalidValueObjectError)
		assert.Equal(t, "Consentimento LGPD", invalidErr.Key)
	})

	t.Run("it should not be able to register when email already exists", func(t *testing.T) {
		useCase, _ := factories.NewRegisterShiftVeterinaryStubFactory()
		consent := time.Now()

		email, _ := sharedvalueobjects.NewEmail("existing@email.com")
		phone1, _ := sharedvalueobjects.NewPhone("11999999999")
		cpf1, _ := sharedvalueobjects.NewCpf("12345678901")
		crmv1, _ := valueobjects.NewCrmv("12345", "SP")
		specialties1, _ := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyGeneralPractice})
		existing, _ := entities.NewShiftVeterinary(
			*email,
			*phone1,
			"password123",
			"Existing Vet",
			*cpf1,
			*crmv1,
			*specialties1,
			*valueobjects.PendingDocumentApproval(),
			&consent,
		)

		_, err := useCase.Execute(&usecases.RegisterShiftVeterinaryUseCaseInput{Veterinary: *existing})
		assert.Nil(t, err)

		phone2, _ := sharedvalueobjects.NewPhone("11888888888")
		cpf2, _ := sharedvalueobjects.NewCpf("98765432100")
		crmv2, _ := valueobjects.NewCrmv("54321", "RJ")
		specialties2, _ := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyGeneralPractice})
		duplicate, _ := entities.NewShiftVeterinary(
			*email,
			*phone2,
			"password123",
			"Duplicate Vet",
			*cpf2,
			*crmv2,
			*specialties2,
			*valueobjects.PendingDocumentApproval(),
			&consent,
		)

		_, err = useCase.Execute(&usecases.RegisterShiftVeterinaryUseCaseInput{Veterinary: *duplicate})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.AlreadyExistsError{}, err)
		alreadyExistsErr := err.(*customerror.AlreadyExistsError)
		assert.Equal(t, "E-mail", alreadyExistsErr.Field)
		assert.Equal(t, "existing@email.com", alreadyExistsErr.Value)
	})
}

