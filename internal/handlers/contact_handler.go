package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LeviVromao/backend_challenge/internal/services"
)

func ContactFormHandler(w http.ResponseWriter, r *http.Request) {
	var form ContactForm
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, "Invaalid Request", http.StatusBadRequest)
		return
	}

	if services.VerifyCaptcha(form.RecaptchaResponse) {
		services.SendMail(form)
		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Captcha Invalido", http.StatusUnauthorized)
	}
}
