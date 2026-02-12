package customerror

import "fmt"

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return fmt.Sprintf("invalid email or password")
}
