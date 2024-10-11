package errors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, title, detail, instance string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		Type:     "about:blank",
		Title:    title,
		Detail:   detail,
		Instance: instance,
	}

	json.NewEncoder(w).Encode(errorResponse)
}
