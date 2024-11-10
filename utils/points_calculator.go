package utils
//Importing packages
import (
    "math"
    "regexp"
    "strings"
    "time"
    "receipt-processor-challenge/models"
)

func CalculatePoints(receipt models.Receipt) int {
    points := 0

    // One point for every alphanumeric character in the retailer name
    alnum := regexp.MustCompile(`[a-zA-Z0-9]`)
    points += len(alnum.FindAllString(receipt.Retailer, -1))

    // 50 points if the total is a round dollar amount with no cents
    if receipt.Total == float64(int(receipt.Total)) {
        points += 50
    }

    // 25 points if the total is a multiple of 0.25
    if math.Mod(receipt.Total, 0.25) == 0 {
        points += 25
    }

    // 5 points for every two items on the receipt
    points += (len(receipt.Items) / 2) * 5

    // If the trimmed length of the item description is a multiple of 3,
    // multiply the price by 0.2 and round up to the nearest integer
    for _, item := range receipt.Items {
        description := strings.TrimSpace(item.ShortDescription)
        if len(description)%3 == 0 {
            itemPoints := int(math.Ceil(item.Price * 0.2))
            points += itemPoints
        }
    }

    // 6 points if the day in the purchase date is odd
    if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
        if date.Day()%2 != 0 {
            points += 6
        }
    }

    // 10 points if the time of purchase is after 2:00pm and before 4:00pm
    if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
        if t.Hour() == 14 {
            points += 10
        }
    }

    return points
}
