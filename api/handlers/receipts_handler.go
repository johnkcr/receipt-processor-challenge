package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/johnkcr/receipt-processor-challenge/api/gen"
	"github.com/johnkcr/receipt-processor-challenge/internal/service"
)

type APIHandler struct {
	receiptService service.ReceiptService
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{
		receiptService: service.NewReceiptService(),
	}
}

// Implements the POST /receipts/process endpoint
func (h *APIHandler) PostReceiptsProcess(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON body into a Receipt object
	var receipt gen.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	// Process the receipt
	id, err := h.receiptService.ProcessReceipt(receipt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process receipt: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the generated receipt ID
	response := gen.ReceiptResponse{Id: id}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

// Implements the GET /receipts/{id}/points endpoint
func (h *APIHandler) GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request, id string) {
	// Retrieve the points for the receipt ID
	points, err := h.receiptService.GetPoints(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve points: %v", err), http.StatusNotFound)
		return
	}

	// Respond with the points
	response := gen.PointsResponse{Points: &points}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
