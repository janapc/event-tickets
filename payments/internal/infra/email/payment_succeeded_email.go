package email

import "fmt"

type PaymentSucceededEmail struct {
	Subject string
	Message string
}

func IntlPaymentSucceeded(language, name string) *PaymentSucceededEmail {
	switch language {
	case "en":
		formatMessage := fmt.Sprintf("Hello %s.\\n Your payment has been approved.", name)
		return &PaymentSucceededEmail{
			Subject: "Payment Approved =)",
			Message: formatMessage,
		}
	case "pt":
		formatMessage := fmt.Sprintf("Olá %s.\\n Seu pagamento foi aprovado.", name)
		return &PaymentSucceededEmail{
			Subject: "Pagamento Aprovado =)",
			Message: formatMessage,
		}
	default:
		formatMessage := fmt.Sprintf("Hello %s.\\n Your payment has been approved.", name)
		return &PaymentSucceededEmail{
			Subject: "Payment Approved =)",
			Message: formatMessage,
		}
	}
}
