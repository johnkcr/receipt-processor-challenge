package main

import (
	"log"
	"net/http"

	"github.com/johnkcr/receipt-processor-challenge/api/gen"
	"github.com/johnkcr/receipt-processor-challenge/api/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Initialize the API handler
	apiHandler := handlers.NewAPIHandler()

	// Create a Chi router
	router := chi.NewRouter()

	// Register the routes using HandlerFromMux
	gen.HandlerFromMux(apiHandler, router)

	// Start the server
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
