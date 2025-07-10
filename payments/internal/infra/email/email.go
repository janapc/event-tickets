package email

import (
	"github.com/go-mail/mail"
	"log"
	"os"
	"strconv"
)

func SendEmail(email string, subject string, message string) error {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_FROM"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	d := mail.NewDialer(os.Getenv("MAIL_HOST"), port, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	log.Println("Email sent")
	return nil
}
