package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"
)

// fetchURL fetches the content of the URL and returns the HTML document.
func fetchURL(url string) (*html.Node, error) {
	// Send an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the HTML response
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// extractTitles traverses the HTML node tree and extracts article titles.
func extractTitles(doc *html.Node) []string {
	var titles []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					titles = append(titles, c.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return titles
}

func main() {
	url := "https://example.com"

	// Fetch the URL
	doc, err := fetchURL(url)
	if err != nil {
		log.Fatalf("Failed to fetch URL: %v", err)
	}

	// Extract and print the titles
	titles := extractTitles(doc)
	fmt.Println("Article Titles:")
	for _, title := range titles {
		fmt.Println(title)
	}
}
