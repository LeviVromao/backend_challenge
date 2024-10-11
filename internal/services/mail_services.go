package services

import (
	"backend_challenge/internal/models"
	"fmt"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendMail(form models.ContactForm) error {

	mailHost := os.Getenv("MAIL_HOST")
	mailPort := 587
	mailUser := os.Getenv("MAIL_AUTH_USER")
	mailName := os.Getenv("MAIL_AUTH_NAME")
	mailPass := os.Getenv("MAIL_AUTH_PASS")

	if mailHost == "" || mailUser == "" || mailPass == "" {
		return fmt.Errorf("missing email environment variables")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mailUser)
	m.SetHeader("To", form.Mail)
	m.SetHeader("Subject", os.Getenv("TEXT_MAIL_TITLE"))

	body := PrepareEmailBody(mailName, mailUser, form.Comment)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(mailHost, mailPort, mailUser, mailPass)
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email successfully sent!")
	return nil
}

func PrepareEmailBody(name string, email string, comment string) string {
	emailTemplate := os.Getenv("TEXT_MAIL_BODY")

	emailTemplate = strings.ReplaceAll(emailTemplate, "{name}", name)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{email}", email)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{comment}", comment)

	return emailTemplate
}

func ValidEmail(email string) bool {
	return len(email) > 5 && len(email) < 254 && strings.Contains(email, "@")
}
