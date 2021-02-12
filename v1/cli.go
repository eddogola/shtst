package main

import (
	"fmt"
	"log"

	"github.com/eddogola/shtst"
)

func main() {
	fmt.Println("Enter the link to the site whose URL you want to shorten")

	var url string
	if _, err := fmt.Scan(&url); err != nil {
		log.Fatalf("Unable to read URL: %v", err)
	}

	// Get the shortened version of the URL
	shURL, err := shtst.GenerateShort(url)
	if err != nil {
		// Get short URL from user (command-line)
		fmt.Println("Enter the shortened URL you'd like")
		fmt.Scan(&shURL)
	}

	urlsh := shtst.URLsh{Original: shtst.SanitizeURL(url), Short: shURL}

	// Run redirect server
	shtst.RedirectToLongURL(urlsh)
}
