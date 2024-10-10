package main

import (
	"log"
	"net/http"

	"github.com/LeviVromao/backend_challenge/internal/routes"
)

func main() {
	r := routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
