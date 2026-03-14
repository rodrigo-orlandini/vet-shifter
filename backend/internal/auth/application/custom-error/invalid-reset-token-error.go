package customerror

type InvalidResetTokenError struct{}

func (e *InvalidResetTokenError) Error() string {
	return "invalid or expired reset token"
}
