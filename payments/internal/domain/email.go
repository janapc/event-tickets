package domain

type IEmail interface {
	sendEmail(email string, subject string, message string) error
}
