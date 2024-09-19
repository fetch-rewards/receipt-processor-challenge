package controller

import (
	"fmt"
	"net/http"
	"receipt-processor-challenge/logic"
	"receipt-processor-challenge/model"

	"github.com/gin-gonic/gin"
)

type ReceiptController struct {
	receiptsService logic.IReceiptsService
}

func NewReceiptsController(receiptsService logic.IReceiptsService) *ReceiptController {
	h := new(ReceiptController)
	h.receiptsService = receiptsService
	return h
}

func (h *ReceiptController) AddRoutes(router *gin.Engine) {
	receipts := router.Group("/receipts")
	receipts.POST("/process", h.ProcessReceipt)
	receipts.GET("/:id/points", h.GetReceiptPoints)
}

// ProcessReceipts stores a Receipt and the points calculated from the receipt.
func (h *ReceiptController) ProcessReceipt(c *gin.Context) {
	// Parse receipt
	var receipt model.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The receipt is invalid",
		})
		return
	}

	// Process receipt
	points := h.receiptsService.ProcessReceipt(&receipt)
	scoredReceipt := model.ReceiptScore{
		Receipt: receipt,
		Points:  points,
	}

	// Store scored receipt
	id, err := h.receiptsService.StoreScoredReceipt(&scoredReceipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The receipt is invalid",
		})
		return
	}

	response := model.IDResponse{ID: id}

	c.JSON(http.StatusOK, response)
}

// GetReceiptPoints returns the points calculated by the id's receipt
func (h *ReceiptController) GetReceiptPoints(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id provided",
		})
	}

	var points int64
	_, err := h.receiptsService.GetByID(id)
	if err == nil {
		points, err = h.receiptsService.GetPointsByID(id)
	}
	if err != nil {
		if err == model.ErrReceiptNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No receipt found for that id",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to retrieve points for receipt id: %v. Error: %v", id, err),
		})
		return
	}

	response := model.PointsResponse{
		Points: points,
	}

	c.JSON(http.StatusOK, response)
}
