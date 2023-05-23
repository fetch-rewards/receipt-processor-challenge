package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func main() {
	// receipts := make(map[string]int)

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

	// Write the same body back to the client.
	w.Write(body)

	jsonFile, error := os.Open("examples/morning-receipt.json")
	fmt.Println(error)
	p, error := ioutil.ReadAll(jsonFile)
	if error != nil {
		fmt.Println(error)
	} else {
		s := string(p)
		fmt.Fprintln(w, s)
	}
	defer jsonFile.Close()
	b := make([]byte, 16)
	_, error2 := rand.Read(b)
	if error2 != nil {
		panic(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Fprintln(w, uuid)
}
