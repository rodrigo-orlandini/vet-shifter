package entities_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
)

func validVetDeps(t *testing.T) (
	sharedvalueobjects.Email,
	sharedvalueobjects.Phone,
	sharedvalueobjects.Cpf,
	valueobjects.Crmv,
	valueobjects.Specialties,
) {
	t.Helper()
	email, _ := sharedvalueobjects.NewEmail("vet@test.com")
	phone, _ := sharedvalueobjects.NewPhone("11999999999")
	cpf, _ := sharedvalueobjects.NewCpf("12345678901")
	crmv, _ := valueobjects.NewCrmv("12345", "SP")
	specialties, _ := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyGeneralPractice})

	return *email, *phone, *cpf, *crmv, *specialties
}

func TestEntityShiftVeterinary(t *testing.T) {
	t.Run("it should create a valid shift veterinary", func(t *testing.T) {
		email, phone, cpf, crmv, specialties := validVetDeps(t)
		consent := time.Now()

		v, err := entities.NewShiftVeterinary(
			email, phone, "password123", "Dr. Vet",
			cpf, crmv, specialties,
			*valueobjects.PendingDocumentApproval(), &consent,
		)

		assert.Nil(t, err)
		assert.NotNil(t, v)
		assert.NotEmpty(t, v.Id)
		assert.Equal(t, "vet@test.com", v.Email.GetValue())
		assert.Equal(t, "11999999999", v.Phone.GetValue())
		assert.Equal(t, "password123", v.Password)
		assert.Equal(t, "Dr. Vet", v.FullName)
		assert.Equal(t, "12345678901", v.Cpf.GetValue())
		assert.Equal(t, "12345", v.Crmv.GetNumber())
		assert.Equal(t, "SP", v.Crmv.GetState())
		assert.Equal(t, 1, v.Specialties.Len())
		assert.NotNil(t, v.ConsentLgpdAt)
		assert.NotNil(t, v.CreatedAt)
	})

	t.Run("it should fail without LGPD consent", func(t *testing.T) {
		email, phone, cpf, crmv, specialties := validVetDeps(t)

		v, err := entities.NewShiftVeterinary(
			email, phone, "password123", "Dr. Vet",
			cpf, crmv, specialties,
			*valueobjects.PendingDocumentApproval(), nil,
		)

		assert.Nil(t, v)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Consentimento LGPD")
	})

	t.Run("it should trim full name", func(t *testing.T) {
		email, phone, cpf, crmv, specialties := validVetDeps(t)
		consent := time.Now()

		v, err := entities.NewShiftVeterinary(
			email, phone, "password123", "  Dr. Vet  ",
			cpf, crmv, specialties,
			*valueobjects.PendingDocumentApproval(), &consent,
		)

		assert.Nil(t, err)
		assert.Equal(t, "Dr. Vet", v.FullName)
	})

	t.Run("it should fail for short full name", func(t *testing.T) {
		email, phone, cpf, crmv, specialties := validVetDeps(t)
		consent := time.Now()

		v, err := entities.NewShiftVeterinary(
			email, phone, "password123", "A",
			cpf, crmv, specialties,
			*valueobjects.PendingDocumentApproval(), &consent,
		)

		assert.Nil(t, v)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Nome completo")
	})

	t.Run("it should fail for empty full name", func(t *testing.T) {
		email, phone, cpf, crmv, specialties := validVetDeps(t)
		consent := time.Now()

		v, err := entities.NewShiftVeterinary(
			email, phone, "password123", "",
			cpf, crmv, specialties,
			*valueobjects.PendingDocumentApproval(), &consent,
		)

		assert.Nil(t, v)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Nome completo")
	})

	t.Run("it should fail for short password", func(t *testing.T) {
		email, phone, cpf, crmv, specialties := validVetDeps(t)
		consent := time.Now()

		v, err := entities.NewShiftVeterinary(
			email, phone, "short", "Dr. Vet",
			cpf, crmv, specialties,
			*valueobjects.PendingDocumentApproval(), &consent,
		)

		assert.Nil(t, v)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Senha")
	})

	t.Run("it should fail for whitespace-only full name", func(t *testing.T) {
		email, phone, cpf, crmv, specialties := validVetDeps(t)
		consent := time.Now()

		v, err := entities.NewShiftVeterinary(
			email, phone, "password123", "   ",
			cpf, crmv, specialties,
			*valueobjects.PendingDocumentApproval(), &consent,
		)

		assert.Nil(t, v)
		assert.NotNil(t, err)
	})
}
