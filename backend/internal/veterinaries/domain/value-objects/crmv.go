package valueobjects

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

type Crmv struct {
	number string
	state  string
}

func NewCrmv(number string, state string) (*Crmv, error) {
	if number == "" {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "CrmvNumber",
			Value: number,
		}
	}

	if len(state) != 2 {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "CrmvState",
			Value: state,
		}
	}

	return &Crmv{number: number, state: state}, nil
}

func (c *Crmv) GetNumber() string {
	return c.number
}

func (c *Crmv) GetState() string {
	return c.state
}
