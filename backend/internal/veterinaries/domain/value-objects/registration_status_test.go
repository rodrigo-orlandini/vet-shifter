package valueobjects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
)

func TestValueObjectRegistrationStatus(t *testing.T) {
	t.Run("it should create pending_document_approval status", func(t *testing.T) {
		status, err := valueobjects.NewRegistrationStatus(valueobjects.RegistrationStatusPendingDocumentApproval)
		assert.Nil(t, err)
		assert.NotNil(t, status)
		assert.Equal(t, valueobjects.RegistrationStatusPendingDocumentApproval, status.String())
	})

	t.Run("it should create complete status", func(t *testing.T) {
		status, err := valueobjects.NewRegistrationStatus(valueobjects.RegistrationStatusComplete)
		assert.Nil(t, err)
		assert.NotNil(t, status)
		assert.Equal(t, valueobjects.RegistrationStatusComplete, status.String())
	})

	t.Run("it should fail for invalid status", func(t *testing.T) {
		status, err := valueobjects.NewRegistrationStatus("unknown")
		assert.Nil(t, status)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Situação do cadastro")
	})

	t.Run("it should fail for empty status", func(t *testing.T) {
		status, err := valueobjects.NewRegistrationStatus("")
		assert.Nil(t, status)
		assert.NotNil(t, err)
	})

	t.Run("PendingDocumentApproval helper should return correct value", func(t *testing.T) {
		status := valueobjects.PendingDocumentApproval()
		assert.Equal(t, valueobjects.RegistrationStatusPendingDocumentApproval, status.String())
	})

	t.Run("Complete helper should return correct value", func(t *testing.T) {
		status := valueobjects.Complete()
		assert.Equal(t, valueobjects.RegistrationStatusComplete, status.String())
	})
}
