package handlers
//Importing packages
import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
// Extract variables from the request URL using Gorilla Mux
// 'vars' is a map of URL parameters, and here we're grabbing the 'id' parameter
    vars := mux.Vars(r)
    id := vars["id"]

    // Check if the receipt ID exists in the 'points' map
    if point, exists := points[id]; exists {
        // If the ID is found, set the response content type to JSON
        w.Header().Set("Content-Type", "application/json")

        // Encode and send the response as JSON with the points for the receipt ID
        json.NewEncoder(w).Encode(map[string]int{"points": point})
    } else {
        // If the ID is not found, respond with a 404 Not Found error
        http.Error(w, "Receipt not found", http.StatusNotFound)
    }
}

