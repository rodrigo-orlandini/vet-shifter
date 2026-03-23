package valueobjects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

func TestValueObjectEmail(t *testing.T) {
	t.Run("it should be able to create new email", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test@example.com")
		assert.Nil(t, err)

		assert.Equal(t, email.GetValue(), "test@example.com")
	})

	t.Run("it should be able to create email with dots", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test.user@example.com")
		assert.Nil(t, err)

		assert.Equal(t, email.GetValue(), "test.user@example.com")
	})

	t.Run("it should be able to create email with plus sign", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test+user@example.com")
		assert.Nil(t, err)

		assert.Equal(t, email.GetValue(), "test+user@example.com")
	})

	t.Run("it should be able to create email with hyphens", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test-user@example-domain.com")
		assert.Nil(t, err)

		assert.Equal(t, email.GetValue(), "test-user@example-domain.com")
	})

	t.Run("it should be able to create email with percent sign", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test%user@example.com")
		assert.Nil(t, err)

		assert.Equal(t, email.GetValue(), "test%user@example.com")
	})

	t.Run("it should be able to create email with underscore", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test_user@example.com")
		assert.Nil(t, err)

		assert.Equal(t, email.GetValue(), "test_user@example.com")
	})

	t.Run("it should be able to create email with single character domain", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test@a.com")
		assert.Nil(t, err)

		assert.Equal(t, email.GetValue(), "test@a.com")
	})

	t.Run("it should fail due to missing @ symbol", func(t *testing.T) {
		email, err := valueobjects.NewEmail("testexample.com")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: testexample.com", err.Error())
	})

	t.Run("it should fail due to missing domain", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test@")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: test@", err.Error())
	})

	t.Run("it should fail due to missing TLD", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test@example")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: test@example", err.Error())
	})

	t.Run("it should fail due to invalid TLD (single character)", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test@example.c")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: test@example.c", err.Error())
	})

	t.Run("it should fail due to missing local part", func(t *testing.T) {
		email, err := valueobjects.NewEmail("@example.com")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: @example.com", err.Error())
	})

	t.Run("it should fail due to multiple @ symbols", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test@example@com")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: test@example@com", err.Error())
	})

	t.Run("it should fail due to spaces in email", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test user@example.com")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: test user@example.com", err.Error())
	})

	t.Run("it should fail due to empty string", func(t *testing.T) {
		email, err := valueobjects.NewEmail("")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: ", err.Error())
	})

	t.Run("it should fail due to invalid characters", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test#user@example.com")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		assert.Equal(t, "E-mail inválido: test#user@example.com", err.Error())
	})

	t.Run("it should fail due to domain starting with dot", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test@.example.com")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		if err != nil {
			assert.Equal(t, "E-mail inválido: test@.example.com", err.Error())
		}
	})

	t.Run("it should fail due to domain ending with dot", func(t *testing.T) {
		email, err := valueobjects.NewEmail("test@example.com.")
		assert.Nil(t, email)
		assert.NotNil(t, err)

		if err != nil {
			assert.Equal(t, "E-mail inválido: test@example.com.", err.Error())
		}
	})
}
