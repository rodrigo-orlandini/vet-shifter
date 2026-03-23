package customerror

type InvalidResetTokenError struct{}

func (e *InvalidResetTokenError) Error() string {
	return "Link inválido ou expirado. Solicite uma nova redefinição de senha."
}
