package main

import (
	"fmt"
	"log"

	"gimingo/shink-api-tests/endpoints"
)

func main() {
	token := endpoints.GetAuthToken()
	hash, err := endpoints.BookShinkHash(token)
	if err != nil {
		log.Fatalf("Error booking shink hash: %v", err)
	}
  
  fmt.Printf("Hash booked: %s\n", hash)
}
