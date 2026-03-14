package repositories

import (
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	veterinariesentities "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
	veterinariesvalueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
)

type StubShiftVeterinaryRepository struct {
	usersByEmail map[string]*veterinariesentities.ShiftVeterinary
	usersByID    map[string]*veterinariesentities.ShiftVeterinary
}

func NewStubShiftVeterinaryRepository() *StubShiftVeterinaryRepository {
	return &StubShiftVeterinaryRepository{
		usersByEmail: make(map[string]*veterinariesentities.ShiftVeterinary),
		usersByID:    make(map[string]*veterinariesentities.ShiftVeterinary),
	}
}

func (r *StubShiftVeterinaryRepository) AddUser(
	email string,
	userID string,
	userType string,
	passwordHash string,
) {
	if userType != sharedvalueobjects.UserTypeShiftVeterinary {
		return
	}

	emailVO, _ := sharedvalueobjects.NewEmail(email)
	phoneVO, _ := sharedvalueobjects.NewPhone("11999999999")
	cpfVO, _ := sharedvalueobjects.NewCpf("12345678901")
	crmvVO, _ := veterinariesvalueobjects.NewCrmv("12345", "SP")
	specialtiesVO, _ := veterinariesvalueobjects.NewSpecialties([]string{veterinariesvalueobjects.SpecialtyGeneralPractice})

	u, err := veterinariesentities.NewShiftVeterinary(
		*emailVO,
		*phoneVO,
		passwordHash,
		"Test Vet",
		*cpfVO,
		*crmvVO,
		*specialtiesVO,
		*veterinariesvalueobjects.Complete(),
		nil,
	)

	if err != nil {
		panic(err)
	}

	u.Id = userID
	r.usersByEmail[email] = u
	r.usersByID[userID] = u
}

func (r *StubShiftVeterinaryRepository) FindByEmail(
	email sharedvalueobjects.Email,
) (*veterinariesentities.ShiftVeterinary, error) {
	if u, ok := r.usersByEmail[email.GetValue()]; ok {
		return u, nil
	}

	return nil, nil
}

func (r *StubShiftVeterinaryRepository) UpdatePassword(userID string, hashedPassword string) error {
	if u, ok := r.usersByID[userID]; ok {
		u.Password = hashedPassword
		return nil
	}

	return nil
}
