package entities

import (
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"time"

	"github.com/google/uuid"
)

const (
	CompanyApprovalPendingDocumentApproval = "pending_document_approval"
	CompanyApprovalComplete                = "complete"
)

type Company struct {
	Id             string
	Cnpj           valueobjects.Cnpj
	Name           string
	ApprovalStatus string
	CreatedAt      *time.Time
}

func NewCompany(cnpj valueobjects.Cnpj, name string) (*Company, error) {
	now := time.Now()
	id, _ := uuid.NewV7()

	return &Company{
		Id:             id.String(),
		Cnpj:           cnpj,
		Name:           name,
		ApprovalStatus: CompanyApprovalPendingDocumentApproval,
		CreatedAt:      &now,
	}, nil
}
