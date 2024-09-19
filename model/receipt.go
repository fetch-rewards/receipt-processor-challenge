package model

import (
	"errors"

	"github.com/shopspring/decimal"
)

var (
	ErrReceiptNotFound = errors.New("this receipt does not exist in the system")
)

// Receipt represents a receipt from a retailer.
type Receipt struct {
	RetailerName string          `json:"retailer" binding:"required"`
	PurchaseDate Date            `json:"purchaseDate" binding:"required"`
	PurchaseTime MilitaryTime    `json:"purchaseTime" binding:"required"`
	Total        decimal.Decimal `json:"total" binding:"required"`
	Items        []Item          `json:"items" binding:"required,min=1"`
}

// Item represents an item on a Receipt's list of Items
type Item struct {
	ShortDescription string          `json:"shortDescription" binding:"required"`
	Price            decimal.Decimal `json:"price" binding:"required"`
}

// ReceiptScore represents a processed receipt and its number of points.
type ReceiptScore struct {
	Receipt Receipt `json:"receipt"`
	Points  int64   `json:"points"`
}

// IDResponse is the response object for a Receipt ID
type IDResponse struct {
	ID string `json:"id"`
}

// PointsResponse is the response object for a Receipt's points
type PointsResponse struct {
	Points int64 `json:"points"`
}
