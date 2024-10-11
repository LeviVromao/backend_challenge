package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type CaptchaResponse struct {
	Success bool `json:"success"`
}

func VerifyCaptcha(recaptchaResponse string) bool {
	recaptchaSecret := os.Getenv("RECAPTCHA_KEY")
	recaptchaURL := os.Getenv("RECAPTCHA_URL")

	data := url.Values{
		"secret":   {recaptchaSecret},
		"response": {recaptchaResponse},
	}

	resp, err := http.PostForm(recaptchaURL, data)
	if err != nil {
		fmt.Println("Erro ao conectar ao reCAPTCHA:", err)
		return false
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Erro ao decodificar a resposta do reCAPTCHA:", err)
		return false
	}

	fmt.Printf("Resposta do reCAPTCHA: %+v\n", result)

	return result["success"].(bool)
}
