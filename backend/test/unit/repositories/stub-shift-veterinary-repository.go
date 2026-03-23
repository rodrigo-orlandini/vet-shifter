package repositories

import (
	"time"

	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	veterinariesentities "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
	veterinariesvalueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
)

type StubShiftVeterinaryRepository struct {
	usersByEmail map[string]*veterinariesentities.ShiftVeterinary
	usersByCpf   map[string]*veterinariesentities.ShiftVeterinary
	usersByPhone map[string]*veterinariesentities.ShiftVeterinary
	usersByID    map[string]*veterinariesentities.ShiftVeterinary
}

func NewStubShiftVeterinaryRepository() *StubShiftVeterinaryRepository {
	return &StubShiftVeterinaryRepository{
		usersByEmail: make(map[string]*veterinariesentities.ShiftVeterinary),
		usersByCpf:   make(map[string]*veterinariesentities.ShiftVeterinary),
		usersByPhone: make(map[string]*veterinariesentities.ShiftVeterinary),
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

	consent := time.Now()
	u, err := veterinariesentities.NewShiftVeterinary(
		*emailVO,
		*phoneVO,
		passwordHash,
		"Test Vet",
		*cpfVO,
		*crmvVO,
		*specialtiesVO,
		*veterinariesvalueobjects.Complete(),
		&consent,
	)

	if err != nil {
		panic(err)
	}

	u.Id = userID
	r.usersByEmail[email] = u
	r.usersByCpf[u.Cpf.GetValue()] = u
	r.usersByPhone[u.Phone.GetValue()] = u
	r.usersByID[userID] = u
}

func (r *StubShiftVeterinaryRepository) Create(
	veterinary veterinariesentities.ShiftVeterinary,
) (*veterinariesentities.ShiftVeterinary, error) {
	v := veterinary

	r.usersByEmail[v.Email.GetValue()] = &v
	r.usersByCpf[v.Cpf.GetValue()] = &v
	r.usersByPhone[v.Phone.GetValue()] = &v
	r.usersByID[v.Id] = &v

	return &v, nil
}

func (r *StubShiftVeterinaryRepository) FindByCpf(
	cpf sharedvalueobjects.Cpf,
) (*veterinariesentities.ShiftVeterinary, error) {
	if u, ok := r.usersByCpf[cpf.GetValue()]; ok {
		return u, nil
	}

	return nil, nil
}

func (r *StubShiftVeterinaryRepository) FindByEmail(
	email sharedvalueobjects.Email,
) (*veterinariesentities.ShiftVeterinary, error) {
	if u, ok := r.usersByEmail[email.GetValue()]; ok {
		return u, nil
	}

	return nil, nil
}

func (r *StubShiftVeterinaryRepository) FindByPhone(
	phone sharedvalueobjects.Phone,
) (*veterinariesentities.ShiftVeterinary, error) {
	if u, ok := r.usersByPhone[phone.GetValue()]; ok {
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
