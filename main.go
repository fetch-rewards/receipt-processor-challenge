package main
//Importing packages
import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "receipt-processor-challenge/handlers"
    
)

func main() {
    r := mux.NewRouter() // Create a new router using the Gorilla Mux package

	// Define a route for processing receipts
    // Path: /receipts/process with HTTP Method: POST
    // Handler Function: ProcessReceipt (in the handlers package)
    r.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST") 

	// Define a route for retrieving points associated with a specific receipt:
    // Path: /receipts/{id}/points, where {id} is a placeholder for the receipt ID with HTTP Method: GET
    r.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	// Start the HTTP server on port 8080 and use the router 'r' to handle incoming requests
    // If there is an error (e.g., port already in use), log.Fatal will log the error and stop the program
    log.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))

}
