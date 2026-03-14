package valueobjects

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

const (
	RegistrationStatusPendingDocumentApproval = "pending_document_approval"
	RegistrationStatusComplete                = "complete"
)

type RegistrationStatus struct {
	value string
}

func NewRegistrationStatus(value string) (*RegistrationStatus, error) {
	if value != RegistrationStatusPendingDocumentApproval && value != RegistrationStatusComplete {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "RegistrationStatus",
			Value: value,
		}
	}

	return &RegistrationStatus{value: value}, nil
}

func (r *RegistrationStatus) String() string {
	return r.value
}

func PendingDocumentApproval() *RegistrationStatus {
	return &RegistrationStatus{value: RegistrationStatusPendingDocumentApproval}
}

func Complete() *RegistrationStatus {
	return &RegistrationStatus{value: RegistrationStatusComplete}
}
