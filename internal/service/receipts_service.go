package service

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"fetch_test/api/gen"
)

type ReceiptService interface {
	ProcessReceipt(receipt gen.Receipt) (string, error)
	GetPoints(receiptID string) (int64, error)
}

type receiptServiceImpl struct {
	receipts map[string]gen.Receipt // In-memory storage for receipts
	points   map[string]int64       // In-memory storage for points
}

func NewReceiptService() ReceiptService {
	return &receiptServiceImpl{
		receipts: make(map[string]gen.Receipt),
		points:   make(map[string]int64),
	}
}

func (s *receiptServiceImpl) ProcessReceipt(receipt gen.Receipt) (string, error) {
	// Generate a unique ID for the receipt
	receiptID := generateID() // You can replace this with a UUID generator

	// Store the receipt
	s.receipts[receiptID] = receipt

	// Calculate points
	points, err := calculatePoints(receipt)
	if err != nil {
		return "", err
	}
	s.points[receiptID] = points

	return receiptID, nil
}

func (s *receiptServiceImpl) GetPoints(receiptID string) (int64, error) {
	points, exists := s.points[receiptID]
	if !exists {
		return 0, errors.New("receipt not found")
	}
	return points, nil
}

// calculatePoints computes the total points for a receipt
func calculatePoints(receipt gen.Receipt) (int64, error) {
	var totalPoints int64

	// Rule 1: One point for every alphanumeric character in the retailer name
	totalPoints += int64(countAlphanumeric(receipt.Retailer))

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid total: %w", err)
	}
	if math.Mod(total, 1) == 0 {
		totalPoints += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		totalPoints += 25
	}

	// Rule 4: 5 points for every two items
	totalPoints += int64((len(receipt.Items) / 2) * 5)

	// Rule 5: Points for item description length and price
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid item price: %w", err)
			}
			totalPoints += int64(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 5 points if the total is greater than 10.00
	if total > 10.0 {
		totalPoints += 5
	}

	// Rule 7: 6 points if the day in the purchase date is odd

	purchaseDate := receipt.PurchaseDate.Time

	if purchaseDate.Day()%2 != 0 {
		totalPoints += 6
	}

	// Rule 8: 10 points if the time of purchase is between 2:00pm and 4:00pm
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, fmt.Errorf("invalid purchase time: %w", err)
	}
	if purchaseTime.Hour() == 14 {
		totalPoints += 10
	}

	return totalPoints, nil
}

// countAlphanumeric counts the number of alphanumeric characters in a string
func countAlphanumeric(s string) int {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(re.FindAllString(s, -1))
}

// generateID generates a simple unique ID (replace with UUID in production)
func generateID() string {
	return fmt.Sprintf("%x", time.Now().UnixNano())
}
