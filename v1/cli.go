package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter the link to the site whose URL you want to shorten")

	var url string
	if _, err := fmt.Scan(&url); err != nil {
		log.Fatalf("Unable to read URL: %v", err)
	}
}