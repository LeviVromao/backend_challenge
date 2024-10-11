package test

// Valida o comportamento da funcao ContactFormHandler qnd recebe uma invalid request
import (
	"backend_challenge/internal/errors"
	"backend_challenge/internal/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContactFormHandler_InvalidRequest(t *testing.T) {
	invalidBody := []byte(`{"name": "", "mail": "invalid-email", "comment": ""}`)
	req, _ := http.NewRequest("POST", "/contact", bytes.NewBuffer(invalidBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.ContactFormHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %v, got %v", http.StatusBadRequest, rr.Code)
	}

	var errResp errors.ErrorResponse
	err := json.NewDecoder(rr.Body).Decode(&errResp)
	if err != nil {
		t.Errorf("failed to decode error response: %v", err)
	}

	if errResp.Title != "BadRequestError" {
		t.Errorf("expected error title 'BadRequestError', got '%s'", errResp.Title)
	}

	if errResp.Detail != "The name is empty" {
		t.Errorf("expected error detail 'The name is empty', got '%s'", errResp.Detail)
	}
}
