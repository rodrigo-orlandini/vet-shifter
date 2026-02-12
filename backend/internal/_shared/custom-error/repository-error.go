package customerror

import "fmt"

type RepositoryError struct {
	Entity string
	Field  string
	Err    error
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("Repository error for field '%s' of entity '%s': %s", e.Field, e.Entity, e.Err.Error())
}
