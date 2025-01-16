package main

import (
	"log"
	"net/http"

	"fetch_test/api/gen"
	"fetch_test/api/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Initialize the router
	router := chi.NewRouter()

	// Set up the handler
	apiHandler := handlers.NewAPIHandler()

	// Register the generated routes
	gen.RegisterHandlers(router, apiHandler)

	// Start the server
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
