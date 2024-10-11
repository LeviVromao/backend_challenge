package routes

import (
	"backend_challenge/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/contact", handlers.ContactFormHandler).Methods("POST", "OPTIONS")
	return r
}
