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
	Street         string
	Number         string
	City           string
	State          string
	ZipCode        string
	ApprovalStatus string
	CreatedAt      *time.Time
}

func NewCompany(cnpj valueobjects.Cnpj, name string, address *Address) (*Company, error) {
	now := time.Now()
	id, _ := uuid.NewV7()
	c := &Company{
		Id:             id.String(),
		Cnpj:           cnpj,
		Name:           name,
		ApprovalStatus: CompanyApprovalPendingDocumentApproval,
		CreatedAt:      &now,
	}
	if address != nil {
		c.Street = address.Street
		c.Number = address.Number
		c.City = address.City
		c.State = address.State
		c.ZipCode = address.ZipCode
	}
	return c, nil
}

type Address struct {
	Street  string
	Number  string
	City    string
	State   string
	ZipCode string
}
