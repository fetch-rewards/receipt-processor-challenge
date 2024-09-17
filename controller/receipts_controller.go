package controller

import (
	"net/http"
	"receipt-processor-challenge/logic"

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
	receipts.POST("/process", h.ProcessReceipts)
	receipts.POST("/:id/points", h.ProcessReceiptPoints)
}

// ProcessReceipts TODO:
func (h *ReceiptController) ProcessReceipts(c *gin.Context) {
	c.JSON(http.StatusOK, "receipts processed")
}

// ProcessReceiptPoints TODO:
func (h *ReceiptController) ProcessReceiptPoints(c *gin.Context) {
	c.JSON(http.StatusOK, "receipt points processed")
}
