package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/utils"

	"github.com/google/uuid"
)

var receipts = make(map[string]models.Receipt)
var points = make(map[string]int)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt

    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id := uuid.New().String()
    receipts[id] = receipt
    points[id] = utils.CalculatePoints(receipt)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"id": id})
}
