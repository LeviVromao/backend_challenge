package handlers

import (
	"backend_challenge/internal/errors"
	"backend_challenge/internal/models"
	"backend_challenge/internal/services"
	"encoding/json"
	"net/http"
)

func ContactFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var form models.ContactForm
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		errors.SendErrorResponse(w, http.StatusBadRequest, "BadRequestError", "Invalid Request Body", "/api-endpoint")
		return
	}

	if form.Name == "" {
		errors.SendErrorResponse(w, http.StatusBadRequest, "BadRequestError", "The name is empty", "/api-endpoint")
		return
	}

	if !services.ValidEmail(form.Mail) {
		errors.SendErrorResponse(w, http.StatusBadRequest, "BadRequestError", "The email is invalid", "/api-endpoint")
		return
	}

	if !services.VerifyCaptcha(form.RecaptchaResponse) {
		errors.SendErrorResponse(w, http.StatusUnauthorized, "UnauthorizedError", "The captcha is incorrect!", "/api-endpoint")
		return
	}

	if err := services.SendMail(form); err != nil {
		errors.SendErrorResponse(w, http.StatusInternalServerError, "InternalServerError", "Internal server error", "/api-endpoint")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
