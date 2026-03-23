package email

import (
	"fmt"

	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	valueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/auth/application/services"

	"github.com/resend/resend-go/v3"
)

type ResendSender struct {
	client *resend.Client
	from   string
}

func NewResendSender() *ResendSender {
	apiKey := utils.GetEmailSenderAPIKey()
	if apiKey == "" {
		return &ResendSender{}
	}

	from := utils.GetEmailSenderFromEmail()

	return &ResendSender{
		client: resend.NewClient(apiKey),
		from:   from,
	}
}

func (s *ResendSender) SendPasswordResetEmail(to valueobjects.Email, resetLink string) error {
	if s.client == nil {
		return fmt.Errorf("EMAIL_SENDER_API_KEY não está definida")
	}

	recipient := to.GetValue()
	if testDest := utils.GetEmailSenderTestDestination(); testDest != "" {
		recipient = testDest
	}

	subject := "Redefinição de senha - Vet Shifter"
	html := fmt.Sprintf(
		`<p>Você solicitou a redefinição de senha.</p>
			<p><a href="%s">Clique aqui para redefinir sua senha</a></p>
			<p>Este link expira em 1 hora. Se você não solicitou, ignore este e-mail.</p>`,
		resetLink,
	)

	_, err := s.client.Emails.Send(&resend.SendEmailRequest{
		From:    s.from,
		To:      []string{recipient},
		Subject: subject,
		Html:    html,
	})
	return err
}

var _ services.EmailSender = (*ResendSender)(nil)
