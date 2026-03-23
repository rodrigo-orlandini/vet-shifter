package valueobjects

import (
	"slices"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

const (
	SpecialtyGeneralPractice  = "general_practice" // clinico geral
	SpecialtyFelines          = "felines"          // felinos
	SpecialtyWildlife         = "wildlife"         // silvestres
	SpecialtyDermatology      = "dermatology"      // dermatologia
	SpecialtyCardiology       = "cardiology"       // cardiologia
	SpecialtyNephrology       = "nephrology"       // nefrologia
	SpecialtyUrology          = "urology"          // urologia
	SpecialtyEndocrinology    = "endocrinology"    // endocrinologia
	SpecialtyGastroenterology = "gastroenterology" // gastroenterologia
	SpecialtyNeurology        = "neurology"        // neurologia
	SpecialtyOrthopedics      = "orthopedics"      // ortopedia
	SpecialtyDentistry        = "dentistry"        // odontologia
	SpecialtyOphthalmology    = "ophthalmology"    // oftalmologia
	SpecialtyUltrasound       = "ultrasound"       // ultrassom
	SpecialtyPathology        = "pathology"        // patologia
	SpecialtyAnesthesiology   = "anesthesiology"   // anestesiologia
	SpecialtyICU              = "icu"              // uti
	SpecialtyOncology         = "oncology"         // oncologia
	SpecialtyPhysiotherapy    = "physiotherapy"    // fisioterapia
	SpecialtyBehavioral       = "behavioral"       // comportamental
)

var allowedSpecialties = []string{
	SpecialtyGeneralPractice,
	SpecialtyFelines,
	SpecialtyWildlife,
	SpecialtyDermatology,
	SpecialtyCardiology,
	SpecialtyNephrology,
	SpecialtyUrology,
	SpecialtyEndocrinology,
	SpecialtyGastroenterology,
	SpecialtyNeurology,
	SpecialtyOrthopedics,
	SpecialtyDentistry,
	SpecialtyOphthalmology,
	SpecialtyUltrasound,
	SpecialtyPathology,
	SpecialtyAnesthesiology,
	SpecialtyICU,
	SpecialtyOncology,
	SpecialtyPhysiotherapy,
	SpecialtyBehavioral,
}

func AllAvailableSpecialties() []string {
	out := make([]string, len(allowedSpecialties))
	copy(out, allowedSpecialties)

	return out
}

type Specialties struct {
	items []string
}

func NewSpecialties(items []string) (*Specialties, error) {
	if items == nil {
		items = []string{}
	}

	list := make([]string, 0, len(items))
	for _, s := range items {
		if !slices.Contains(allowedSpecialties, s) {
			return nil, &customerror.InvalidValueObjectError{
				Key:   "Especialidades",
				Value: s,
			}
		}

		list = append(list, s)
	}

	if len(list) < 1 {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Especialidades",
			Value: "selecione ao menos uma especialidade",
		}
	}

	return &Specialties{items: list}, nil
}

func (s *Specialties) GetValue() []string {
	if s == nil || s.items == nil {
		return nil
	}

	out := make([]string, len(s.items))
	copy(out, s.items)

	return out
}

func (s *Specialties) Len() int {
	if s == nil {
		return 0
	}

	return len(s.items)
}
