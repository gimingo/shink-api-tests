package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"gimingo/shink-api-tests/endpoints"
)

func bookAndCreateHashUntilQuoteLimitReached(token string) {
	quotaLimit := 100
	bookedHashes := 0
	createdShinks := 0

	for i := 1; i <= quotaLimit; i++ {
		hash, err := endpoints.BookShinkHash(token)
		if err != nil {
			log.Fatalf("Iteration %d: Error booking shink hash: %v", i, err)
		}
		bookedHashes++

		_, err = endpoints.CreateShink(token, "shink-"+strconv.Itoa(i), "https://example.com", hash)
		if err != nil {
			log.Printf("Iteration %d: Error creating shink: %v", i, err)
		} else {
			createdShinks++
		}

		printProgressBar(i, quotaLimit)
	}

	fmt.Println()
	log.Printf("Successfully booked %d hashes and created %d shinks!", bookedHashes, createdShinks)
}

func main() {
	token := endpoints.GetAuthToken()
	bookAndCreateHashUntilQuoteLimitReached(token)
}

func printProgressBar(current, total int) {
	width := 50
	percentage := float64(current) / float64(total)
	filled := int(percentage * float64(width))
	bar := "[" + strings.Repeat("#", filled) + strings.Repeat("-", width-filled) + "]"
	fmt.Printf("\rProgress: %s %d/%d", bar, current, total)
}

