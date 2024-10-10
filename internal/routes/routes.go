package routes

import (
	"github.com/LeviVromao/backend_challenge/internal/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/contact", handlers.ContactFormHandler).Methods("POST")
	return r
}
