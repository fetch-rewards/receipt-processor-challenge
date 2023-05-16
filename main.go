package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []item `json:"items"`
}

var receipts []receipt

func getReciepts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func createReceipt(c *gin.Context) {
	var newReceipt receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	newReceipt.ID = strconv.Itoa(len(receipts) + 1)

	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}

func main() {
	router := gin.Default()
	router.GET("/receipts", getReciepts)
	router.POST("/receipts", createReceipt)
	router.Run("localhost:8080")
}
