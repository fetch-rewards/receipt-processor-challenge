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
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\nMy name is Imad")
	})

	http.HandleFunc("/receipts/process", generateID)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func generateID(w http.ResponseWriter, r *http.Request) {
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
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Fprintln(w, uuid)
}
