package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/gorilla/mux"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Points struct {
	Points int `json:"points"`
}

type Id struct {
	Id string `json:"id"`
}

var receipts = make(map[string]int)
var mu = &sync.RWMutex{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\nMy name is Imad")
	})

	r.HandleFunc("/receipts/process", generateID)

	r.HandleFunc("/receipts/{id}/points", fetchPointsByID)

	http.Handle("/", r)

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
	} else {
		uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

		points, err := CalculatePoints(bytes.NewReader(body))
		if err != nil {
			http.Error(w, "Error calculating points", http.StatusInternalServerError)
			return
		}

		mu.Lock()
		receipts[uuid] = points
		mu.Unlock()

		err = json.NewEncoder(w).Encode(Id{Id: uuid})
		if err != nil {
			http.Error(w, "Error encoding response body", http.StatusInternalServerError)
			return
		}

	}
}

func CalculatePoints(body io.Reader) (int, error) {
	var re Receipt
	err := json.NewDecoder(body).Decode(&re)
	if err != nil {
		return 0, fmt.Errorf("error decoding JSON: %w", err)
	}

	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	for _, char := range re.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	// Convert total to float64 for calculations.
	total, err := strconv.ParseFloat(re.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing total: %w", err)
	}

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	if total == math.Round(total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if math.Remainder(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += (len(re.Items) / 2) * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range re.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			itemPrice, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, fmt.Errorf("error parsing item price: %w", err)
			}
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	dateParts := strings.Split(re.PurchaseDate, "-")
	if day, err := strconv.Atoi(dateParts[2]); err == nil && day%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	timeParts := strings.Split(re.PurchaseTime, ":")
	if hour, err := strconv.Atoi(timeParts[0]); err == nil && (hour >= 14 && hour < 16) {
		points += 10
	}

	return points, nil
}

func fetchPointsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	mu.RLock()
	points, ok := receipts[id]
	mu.RUnlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "No points found for matching ID",
		})
		return
	}

	err := json.NewEncoder(w).Encode(Points{Points: points})
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}
}
