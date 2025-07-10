package payment

const FailedEventName = "PAYMENT_FAILED"

type FailedEvent struct {
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserLanguage string `json:"user_language"`
}

func (e FailedEvent) Name() string {
	return FailedEventName
}
