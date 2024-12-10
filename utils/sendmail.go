package utils

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", "your-email@example.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)
	dialer := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-password")
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
