package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

type Points struct {
	Points int `json:"points"`
}

var receipts = make(map[string]int)
var mu = &sync.RWMutex{}

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\nMy name is Imad")
	})

	http.HandleFunc("/receipts/process", generateID)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func generateID(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		http.Error(w, "Error generating UUID", http.StatusInternalServerError)
		return
	}

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	points, err := CalculatePoints(bytes.NewReader(body))
	if err != nil {
		http.Error(w, "Error calculating points", http.StatusInternalServerError)
		return
	}

	mu.Lock()
	receipts[uuid] = points
	mu.Unlock()

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the points struct into JSON and send it in the response
	err = json.NewEncoder(w).Encode(Points{Points: points})
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}
}

func CalculatePoints(body io.Reader) (int, error) {
	// var re Receipt
	// err := json.NewDecoder(body).Decode(&re)
	// if err != nil {
	// 	return 0, fmt.Errorf("error decoding JSON: %w", err)
	// }

	points := 10 * 10

	return points, nil
}
