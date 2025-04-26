package store

import "go-based-payment-tracking-api/models"

// In-memory store for payments
var Payments = make(map[string]models.Payment)
