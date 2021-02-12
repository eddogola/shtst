package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

type URLsh struct {
	original string
	short string
}

var urls []URLsh

// SaveLink saves a, generated or provided, short URL
func SaveLink(link string) bool {
	return true
}

// Utility function to check if a string elements is in a string slice
func contains(words []string, word string) bool {
	for _, val := range words {
		if val == word {
			return true
		}
	}
	
	return false
}

// Utility function to remove a string element from a string slice
func remove(words []string, word string) []string {
	for i, w := range words {
		if w == word {
			return append(words[:i], words[i+1:]...)
		}
	}

	return words
}

// GenerateShort takes an original URL and randomly generates a shorter URL, under 30 characters, or an error is raised 
func GenerateShort(link string) (s string, err error) {
	// Parse url
	u, err := url.Parse(link)

	// Get the Host name and path elements
	words := append(strings.Split(u.Host, "."), strings.Split(u.Path, "/")...)

	// Omit common words from link
	omitWords := []string{"www", "com", "org", "net", "site", "gov", "biz"} // Words to eliminate
	for _, omitWord := range omitWords {
		if contains(words, omitWord) {
			words = remove(words, omitWord)
		}
	}

	// Create the short string
	var w string
	for _, word := range words {
		w += word
	}
	
	// Get shortened URL depending on length of string w
	switch count := len(w); {
	case count < 10:
		s = w[:count]
	case count > 10:
		s = w[:10]
	case count > 20:
		s = w[:15]
	case count > 30:
		s = w[:30]
	}

	return
}

func main() {
	fmt.Println("Enter the link to the site whose URL you want to shorten")

	var url string
	if _, err := fmt.Scan(&url); err != nil {
		log.Fatalf("Unable to read URL: %v", err)
	}

	// Get the shortened version of the URL
	shURL, err := GenerateShort(url)
	if err != nil {
		// Get short URL from user (command-line)
		fmt.Println("Enter the shortened URL you'd like")
		fmt.Scan(&shURL)
	}

	urlsh := URLsh{original: url, short: shURL}

	fmt.Println(urlsh)
	
}