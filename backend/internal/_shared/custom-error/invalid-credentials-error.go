package customerror

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "E-mail ou senha incorretos."
}
