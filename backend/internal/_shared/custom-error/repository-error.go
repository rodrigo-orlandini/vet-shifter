package customerror

import "fmt"

type RepositoryError struct {
	Entity string
	Field  string
	Err    error
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("Erro ao buscar %s pelo %s %s", e.Entity, e.Field, e.Err.Error())
}
