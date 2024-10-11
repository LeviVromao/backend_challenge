package main

import (
	"log"
	"net/http"
	"os"

	"backend_challenge/internal/handlers"
	"backend_challenge/pkg/config"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	r := mux.NewRouter()

	r.HandleFunc("/contact", handlers.ContactFormHandler).Methods("POST", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
