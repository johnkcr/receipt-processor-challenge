package handlers

import (
	"context"
	"fmt"

	"fetch_test/api/gen"
	"fetch_test/internal/service"
)

type APIHandler struct {
	receiptService service.ReceiptService
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{
		receiptService: service.NewReceiptService(),
	}
}

// PostReceiptsProcess handles the submission of receipts for processing.
func (h *APIHandler) PostReceiptsProcess(ctx context.Context, receipt gen.Receipt) (gen.ReceiptResponse, error) {
	// Process the receipt using the service layer
	id, err := h.receiptService.ProcessReceipt(receipt)
	if err != nil {
		return gen.ReceiptResponse{}, fmt.Errorf("error processing receipt: %w", err)
	}

	// Return the generated receipt ID
	return gen.ReceiptResponse{ID: id}, nil
}

// GetReceiptsIdPoints retrieves the points awarded for a receipt by ID.
func (h *APIHandler) GetReceiptsIdPoints(ctx context.Context, id string) (gen.PointsResponse, error) {
	// Fetch points for the receipt ID from the service layer
	points, err := h.receiptService.GetPoints(id)
	if err != nil {
		return gen.PointsResponse{}, fmt.Errorf("error retrieving points: %w", err)
	}

	// Return the points in the response
	return gen.PointsResponse{Points: points}, nil
}
