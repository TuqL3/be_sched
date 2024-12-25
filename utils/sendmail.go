package utils

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", "noreply@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "tunglv157@gmail.com", "wfyn ssjb akuf wkrh")
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
