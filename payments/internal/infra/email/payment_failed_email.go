package email

import "fmt"

type PaymentFailedEmail struct {
	Subject string
	Message string
}

func IntlPaymentFailed(language, name string) *PaymentSucceededEmail {
	switch language {
	case "en":
		formatMessage := fmt.Sprintf("Hello %s.\\n Your payment has been rejected, try again.", name)
		return &PaymentSucceededEmail{
			Subject: "Payment Rejected =(",
			Message: formatMessage,
		}
	case "pt":
		formatMessage := fmt.Sprintf("Olá %s.\\n Seu pagamento foi rejeitado, porfavor tente novamente.", name)
		return &PaymentSucceededEmail{
			Subject: "Pagamento Rejeitado =(",
			Message: formatMessage,
		}
	default:
		formatMessage := fmt.Sprintf("Hello %s.\\n Your payment has been rejected, try again.", name)
		return &PaymentSucceededEmail{
			Subject: "Payment Rejected =(",
			Message: formatMessage,
		}
	}
}
