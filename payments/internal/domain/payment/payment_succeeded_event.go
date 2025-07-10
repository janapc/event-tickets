package payment

const SucceededEventName = "PAYMENT_SUCCEEDED"

type SucceededEvent struct {
	UserName         string `json:"user_name"`
	UserEmail        string `json:"user_email"`
	EventId          string `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventImageUrl    string `json:"event_image_url"`
	UserLanguage     string `json:"user_language"`
}

func (e SucceededEvent) Name() string {
	return SucceededEventName
}
