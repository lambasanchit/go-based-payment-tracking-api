package models

// Payment represents a payment object
type Payment struct {
	ID      string  `json:"id"`
	Amount  float64 `json:"amount"`
	Method  string  `json:"method"`
	Status  string  `json:"status"` // eg-> pending, completed, retrying
	Created string  `json:"created"`
}
