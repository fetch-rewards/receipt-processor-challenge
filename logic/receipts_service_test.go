package logic

import (
	"encoding/json"
	"receipt-processor-challenge/model"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReceiptServiceTestSuite struct {
	suite.Suite
	receiptScore model.ReceiptScore
}

func (suite *ReceiptServiceTestSuite) SetupTest() {
	jsonPayload := `
	{
    	"retailer": "Target",
    	"purchaseDate": "2022-01-02",
    	"purchaseTime": "13:13",
    	"total": "1.25",
    	"items": [
    	    {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
    	]
	}`
	err := json.Unmarshal([]byte(jsonPayload), &suite.receiptScore.Receipt)
	if err != nil {
		panic(err)
	}
	suite.receiptScore.Points = 31 // expected
}

// TestProcessReceipt tests the ProcessReceipt function against a predefined payload
func (suite *ReceiptServiceTestSuite) TestProcessReceipt() {
	rs := ReceiptsService{}
	points := rs.ProcessReceipt(&suite.receiptScore.Receipt)
	suite.Equal(suite.receiptScore.Points, points)
}

func TestReceiptServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ReceiptServiceTestSuite))
}
