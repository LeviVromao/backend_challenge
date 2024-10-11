package test

// Testa se o envio do Email vai ser bem sucedido e

// verifica o comportamento da funcao TestSendMail_MissingEnvVariables qnd as variaveis de ambiente nao estao definidas

import (
	"backend_challenge/internal/models"
	"backend_challenge/internal/services"
	"os"
	"testing"
)

func TestSendMail_MissingEnvVariables(t *testing.T) {
	os.Setenv("MAIL_HOST", "")
	form := models.ContactForm{
		Name:    "NAME para teste",
		Mail:    "my.name@example.com",
		Comment: "This is a test comment",
	}

	err := services.SendMail(form)
	if err == nil {
		t.Error("expected an error due to missing MAIL_HOST")
	}
}

func TestSendMail_Success(t *testing.T) {
	os.Setenv("MAIL_HOST", "smtp.gmail.com")
	os.Setenv("MAIL_AUTH_USER", "my.email@gmail.com")
	os.Setenv("MAIL_AUTH_PASS", "passowrd email")

	form := models.ContactForm{
		Name:    "",
		Mail:    "john.doe@example.com",
		Comment: "This is a test comment",
	}

	err := services.SendMail(form)
	if err != nil {
		t.Errorf("failed to send mail: %v", err)
	}
}
