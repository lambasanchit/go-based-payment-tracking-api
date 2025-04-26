package router

import (
	"go-based-payment-tracking-api/handlers"

	"github.com/gorilla/mux"
)

// SetupRouter initializes all the routes
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/payment", handlers.CreatePaymentHandler).Methods("POST")           // Create a new payment.
	r.HandleFunc("/payment/{id}", handlers.GetPaymentStatusHandler).Methods("GET")    // Get the status of an existing payment.
	r.HandleFunc("/payment/retry/{id}", handlers.RetryPaymentHandler).Methods("POST") //  Retry a failed payment by changing its status.

	return r
}
