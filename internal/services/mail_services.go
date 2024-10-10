package services

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

type MailDetails struct {
	Recipient string
	Name      string
	Comment   string
}

func SendMail(details MailDetails) error {
	mailHost := os.Getenv("MAIL_HOST")
	mailPort := 587
	mailUser := os.Getenv("MAIL_AUTH_USER")
	mailPass := os.Getenv("MAIL_AUTH_PASS")

	m := gomail.NewMessage()
	m.SetHeader("from", mailUser)
	m.SetHeader("To", details.Recipient)
	m.SetHeader("Subject", os.Getenv("TEXT_MAIL_TITLE"))

	body := fmt.Sprintf(os.Getenv("TEXT_MAIL_BODY"), details.Name, details.Recipient, details.Comment)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(mailHost, mailPort, mailUser, mailPass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
