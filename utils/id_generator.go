package utils

import "github.com/google/uuid"

// GenerateID generates a unique ID for payments (using UUID)
func GenerateID() string {
	return uuid.New().String()
}
