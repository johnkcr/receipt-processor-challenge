package main

import (
	"log"
	"net/http"
	"os"

	"github.com/johnkcr/receipt-processor-challenge/api/gen"
	"github.com/johnkcr/receipt-processor-challenge/api/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize the API handler
	apiHandler := handlers.NewAPIHandler()

	// Create a Chi router
	router := chi.NewRouter()

	// Register the routes using HandlerFromMux
	gen.HandlerFromMux(apiHandler, router)

	// Start the server
	log.Printf("Server is running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
