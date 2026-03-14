package customerror

import "fmt"

type NotFoundError struct {
	Key   string
	Value string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found '%s' with value: %s", e.Key, e.Value)
}
