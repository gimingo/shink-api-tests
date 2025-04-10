package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"gimingo/shink-api-tests/endpoints"
)

func main() {
	token, err := endpoints.GetAuthToken()
  if err != nil {
    log.Printf("Error creating user") 
  } 

	bookAndCreateHashUntilQuoteLimitReached(token)
}

func bookAndCreateHashUntilQuoteLimitReached(token string) {
  numberOfShinks := 110

	for i := 1; i <= numberOfShinks; i++ {
    bookAndCreateShink(token, i);
		printProgressBar(i, numberOfShinks)
	}

	fmt.Println()
  log.Printf("Successfully booked %d hashes and created %d shinks!", numberOfShinks, numberOfShinks)
}

func bookAndCreateShink(token string, i int) {
		hash, err := endpoints.BookShinkHash(token)
		if err != nil {
			log.Fatalf("Iteration %d: Error booking shink hash: %v", i, err)
		}

		_, err = endpoints.CreateShink(token, "shink-"+strconv.Itoa(i), "https://example.com", hash)
		if err != nil {
			log.Printf("Iteration %d: Error creating shink: %v", i, err)
    }
}

func printProgressBar(current, total int) {
	width := 50
	percentage := float64(current) / float64(total)
	filled := int(percentage * float64(width))
	bar := "[" + strings.Repeat("#", filled) + strings.Repeat("-", width-filled) + "]"
	fmt.Printf("\rProgress: %s %d/%d", bar, current, total)
}

