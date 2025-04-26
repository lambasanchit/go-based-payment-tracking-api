package main

import (
	"fmt"
	"go-based-payment-tracking-api/router"
	"log"
	"net/http"
)

func main() {
	// Use the SetupRouter to initialize the routes
	r := router.SetupRouter()

	// Start the server
	port := ":8080"
	fmt.Println("Server running at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
