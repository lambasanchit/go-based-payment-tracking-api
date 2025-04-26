package handlers

import (
	"encoding/json"
	"go-based-payment-tracking-api/models"
	"go-based-payment-tracking-api/store"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(payment.Method) == "" || payment.Amount <= 0 {
		http.Error(w, "Invalid payment data", http.StatusBadRequest)
		return
	}

	payment.ID = uuid.New().String() // generating uniq id
	store.Payments[payment.ID] = payment

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func GetPaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	payment, exists := store.Payments[id]
	if !exists {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func RetryPaymentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	payment, exists := store.Payments[id]
	if !exists {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	/* Logic for retrying payment (this can be enhanced later)
	Mark as 'retrying' status temporarily (or update logic)
	*/
	payment.Status = "retrying"

	store.Payments[id] = payment

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}
