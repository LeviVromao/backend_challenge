package models

type ContactForm struct {
	Name              string `json:"name"`
	Mail              string `json:"mail"`
	Comment           string `json:"comment"`
	RecaptchaResponse string `json:"g-recaptcha-response"`
}
