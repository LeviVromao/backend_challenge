package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	requiredVars := []string{
		"PORT", "RECAPTCHA_KEY", "RECAPTCHA_URL", "MAIL_HOST",
		"MAIL_PORT", "MAIL_AUTH_USER", "MAIL_AUTH_PASS",
	}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("A variável de ambiente %s não foi definida", v)
		}
	}
}
