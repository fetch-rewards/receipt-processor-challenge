package logic

import (
	"receipt-processor-challenge/adapter/repository"
	"receipt-processor-challenge/model"
	"strings"
	"time"
	"unicode"

	"github.com/shopspring/decimal"
)

type IReceiptsService interface {
	ProcessReceipt(receipt *model.Receipt) (points int64)
	StoreReceipt(receipt *model.Receipt) (id string)
	StoreScoredReceipt(receiptScore *model.ReceiptScore) (id string, err error)
	GetByID(id string) (model.Receipt, error)
	GetPointsByID(id string) (int64, error)
}

type ReceiptsService struct {
	receiptService repository.ReceiptRepo
}

func NewReceiptsService() *ReceiptsService {
	return &ReceiptsService{}
}

func (s *ReceiptsService) ProcessReceipt(receipt *model.Receipt) (points int64) {
	// One point for every alphanumeric character in the retailer name.
	points += s.countAlphanumericCharacters(receipt.RetailerName)

	// 50 points if the total is a round dollar amount with no cents (total == math.ceil(total))
	if receipt.Total.Equal(receipt.Total.Ceil()) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25 (total % 0.25 == 0)
	if receipt.Total.Mod(decimal.NewFromFloat32(0.25)).Equal(decimal.Zero) {
		points += 25
	}

	// 5 points for every two items on the receipt (len(items) / 2 * 5)
	itemCouples := (len(receipt.Items) / 2)
	points += int64(itemCouples * 5)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	var scalar = decimal.NewFromFloat32(0.2)
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += item.Price.Mul(scalar).Ceil().IntPart()
		}
	}

	// 6 points if the day in the purchase date is odd.
	if receipt.PurchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm (14:00) and before 4:00pm (16:00)
	start, _ := time.Parse("15:04", "14:00")
	end, _ := time.Parse("15:04", "16:00")
	if receipt.PurchaseTime.IsBetween(start, end) {
		points += 10
	}

	return
}

func (s *ReceiptsService) StoreReceipt(receipt *model.Receipt) (id string) {
	return s.receiptService.Store(*receipt)
}

func (s *ReceiptsService) StoreScoredReceipt(receiptScore *model.ReceiptScore) (id string, err error) {
	if receiptScore == nil {
		return "", nil
	}

	// tx
	id = s.StoreReceipt(&receiptScore.Receipt)
	err = s.receiptService.StorePoints(id, receiptScore.Points)
	return
}

func (s *ReceiptsService) GetByID(id string) (model.Receipt, error) {
	return s.receiptService.GetByID(id)
}

func (s *ReceiptsService) GetPointsByID(id string) (int64, error) {
	return s.receiptService.GetPointsByID(id)
}

func (s *ReceiptsService) countAlphanumericCharacters(str string) (count int64) {
	for _, char := range str {
		if unicode.IsLetter(char) {
			count++
		}
	}
	return
}
