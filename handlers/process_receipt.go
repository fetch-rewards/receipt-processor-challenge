package handlers
//Importing packages
import (
    "encoding/json"
    "net/http"
    "github.com/google/uuid"
    "receipt-processor-challenge/models"
    "receipt-processor-challenge/utils"
)

// Map to store receipts by their unique ID
var receipts = make(map[string]models.Receipt)

// Map to store points associated with each receipt ID
var points = make(map[string]int)

// ProcessReceipt handles POST requests to process a receipt and store its data
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    // Declare a variable to hold the decoded receipt data
    var receipt models.Receipt

    // Decode the JSON request body into the 'receipt' variable
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        // If there's an error in decoding, send a 400 Bad Request response with the error message
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Generate a unique ID for the receipt using UUID
    id := uuid.New().String()

    // Store the receipt data in the 'receipts' map with the generated ID as the key
    receipts[id] = receipt

    // Calculate the points for the receipt using a utility function and store it in the 'points' map
    points[id] = utils.CalculatePoints(receipt)

    // Set the response content type to JSON
    w.Header().Set("Content-Type", "application/json")

    // Encode and send a JSON response with the generated receipt ID
    json.NewEncoder(w).Encode(map[string]string{"id": id})
}
