// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Receipt struct {
	Retailer     string    `json:"retailer"`
	Total        float64   `json:"total"`
	Items        []string  `json:"items"`
	PurchaseDate time.Time `json:"purchaseDate"`
}

type ResponseID struct {
	ID string `json:"id"`
}

type ResponsePoints struct {
	Points int `json:"points"`
}

var receiptStore map[string]Receipt

func main() {
	receiptStore = make(map[string]Receipt)

	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", ProcessReceipts).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	id := generateID()
	receiptStore[id] = receipt

	response := ResponseID{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	receipt, ok := receiptStore[id]
	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := calculatePoints(receipt)
	response := ResponsePoints{Points: points}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func generateID() string {
	return uuid.New().String()
}

func calculatePoints(receipt Receipt) int {
	points := len(receipt.Retailer)
	if receipt.Total == float64(int(receipt.Total)) {
		points += 50
	}
	if receipt.Total/0.25 == float64(int(receipt.Total/0.25)) {
		points += 25
	}
	points += len(receipt.Items) / 2

	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item))
		if trimmedLength%3 == 0 {
			points += int(receipt.Total * 0.2)
		}
	}

	if receipt.PurchaseDate.Day()%2 == 1 {
		points += 6
	}

	if receipt.PurchaseDate.Hour() >= 14 && receipt.PurchaseDate.Hour() < 16 {
		points += 10
	}

	return points
}
