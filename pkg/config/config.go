package config

import (
	"log"
	"os"
)

func LoadConfig() {
	requiredEnvVars := []string{
		"PORT",
		"ORIGINS",
		"RECAPTCHA_KEY",
		"RECAPTCHA_URL",
		"MAIL_HOST",
		"MAIL_PORT",
		"MAIL_AUTH_USER",
		"MAIL_AUTH_PASS",
		"TEXT_MAIL_TITLE",
		"TEXT_MAIL_BODY",
	}

	for _, v := range requiredEnvVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Variável de ambiente %s não definida", v)
		}
	}
}
