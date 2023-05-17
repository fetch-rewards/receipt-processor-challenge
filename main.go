package main

import (
	"errors"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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
var pointsPerReceiptMap = make(map[string]int)

func getReceipts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func createReceipt(r receipt) receipt {
	r.ID = strconv.Itoa(len(receipts) + 1)

	receipts = append(receipts, r)
	return r
}

func newReceipt(c *gin.Context) {
	var newReceipt receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	newReceipt = createReceipt(newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}

// processes all rules for a given receipt
func processPointsForReceipt(c *gin.Context) {
	totalPoints := 0

	var receiptToProcess receipt

	if err := c.BindJSON(&receiptToProcess); err != nil {
		return
	}

	receiptToProcess = createReceipt(receiptToProcess)

	totalPoints += alphanumericCharactersInRetailer(receiptToProcess.Retailer)
	totalPoints += roundDollarAmount(receiptToProcess.Total)
	totalPoints += multipleOfQuarter(receiptToProcess.Total)
	totalPoints += everyTwoItems(receiptToProcess.Items)
	totalPoints += itemDescription(receiptToProcess.Items)
	totalPoints += isDateOdd(receiptToProcess.PurchaseDate)
	totalPoints += isTimeBetweenFourAndTwo(receiptToProcess.PurchaseTime)

	pointsPerReceiptMap[receiptToProcess.ID] = totalPoints

	c.IndentedJSON(http.StatusOK, gin.H{"id": receiptToProcess.ID})
}

// One point for every alphanumeric character in the retailer name.
func alphanumericCharactersInRetailer(retailerName string) int {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleanedString := nonAlphanumericRegex.ReplaceAllString(retailerName, "")
	return len(cleanedString)
}

// 50 points if the total is a round dollar amount with no cents.
func roundDollarAmount(total string) int {
	if strings.HasSuffix(total, ".00") {
		return 50
	}

	return 0
}

// 25 points if the total is a multiple of 0.25
func multipleOfQuarter(total string) int {
	floatTotal, err := strconv.ParseFloat(total, 32)

	if err != nil {
		return 0
	}

	intTotal := int(floatTotal * 100)

	if intTotal%25 == 0 {
		return 25
	}

	return 0
}

// 5 points for every two items on the receipt.
func everyTwoItems(items []item) int {
	return (len(items) / 2) * 5
}

// If the trimmed length of the item description is a multiple of 3,
// multiply the price by 0.2 and round up to the nearest integer.
// The result is the number of points earned.
func itemDescription(items []item) int {
	points := 0
	for _, item := range items {
		if len(strings.Trim(item.ShortDescription, " "))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 32)
			if err != nil {
				continue
			}
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}

// 6 points if the day in the purchase date is odd.
func isDateOdd(date string) int {
	day, err := strconv.Atoi(strings.Split(date, "-")[2])
	if err != nil {
		return 0
	}
	if day%2 != 0 {
		return 6
	}
	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func isTimeBetweenFourAndTwo(purchaseTime string) int {
	hour, err := strconv.Atoi(strings.Split(purchaseTime, ":")[0])

	if err != nil {
		return 0
	}

	if 14 <= hour && hour < 16 {
		return 10
	}

	return 0
}

func getPointsByReceiptId(id string) (int, error) {

	points, ok := pointsPerReceiptMap[id]
	if ok {
		return points, nil
	}

	return -1, errors.New("no receipt found for that id")
}

func getPointsById(c *gin.Context) {
	id := c.Param("id")

	pointsForReceipt, err := getPointsByReceiptId(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"description": "No receipt found for that id"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"points": pointsForReceipt})
}

func main() {
	router := gin.Default()
	router.GET("/receipts", getReceipts)
	router.POST("/receipts", newReceipt)
	router.POST("/receipts/process", processPointsForReceipt)
	router.GET("/receipts/:id/points", getPointsById)
	router.Run("localhost:8080")
}
