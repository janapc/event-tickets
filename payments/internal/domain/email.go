package domain

import "context"

type IEmail interface {
	sendEmail(ctx context.Context, email string, subject string, message string) error
}
