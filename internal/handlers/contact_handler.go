package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LeviVromao/backend_challenge/internal/services"
)

type ContactForm struct {
	Name              string `json:"name"`
	Mail              string `json:"mail"`
	Comment           string `json:"comment"`
	RecaptchaResponse string `json:"g-recaptcha-response"`
}

func ContactFormHandler(w http.ResponseWriter, r *http.Request) {
	var form ContactForm
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if services.VerifyCaptcha(form.RecaptchaResponse) {
		mailDetails := services.MailDetails{
			Recipient: form.Mail,
			Name:      form.Name,
			Comment:   form.Comment,
		}

		services.SendMail(mailDetails)
		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Captcha Invalido", http.StatusUnauthorized)
	}
}
