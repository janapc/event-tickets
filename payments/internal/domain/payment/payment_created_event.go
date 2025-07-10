package payment

var CreatedEventName = "PAYMENT_CREATED"

type CreatedEvent struct {
	UserName         string  `json:"user_name"`
	UserEmail        string  `json:"user_email"`
	EventId          string  `json:"event_id"`
	EventAmount      float64 `json:"event_amount"`
	PaymentToken     string  `json:"payment_token"`
	EventName        string  `json:"event_name"`
	EventDescription string  `json:"event_description"`
	EventImageUrl    string  `json:"event_image_url"`
	UserLanguage     string  `json:"user_language"`
	PaymentID        string  `json:"payment_id"`
}

func (e CreatedEvent) Name() string {
	return CreatedEventName
}
