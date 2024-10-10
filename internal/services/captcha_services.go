package services

import (
	"encoding/json"
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
		return false
	}
	defer resp.Body.Close()

	var result CaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false
	}

	return result.Success
}
