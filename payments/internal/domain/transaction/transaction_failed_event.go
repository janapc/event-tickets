package transaction

const FailedEventName = "TRANSACTION_FAILED"

type FailedEvent struct {
	PaymentID    string `json:"payment_id"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserLanguage string `json:"user_language"`
}

func (e FailedEvent) Name() string {
	return FailedEventName
}
