package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.airkorea.or.kr/web/vicinityStation?item_code=10008&station_code=111162")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".al2").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("th").Text()
		title := s.Find("td").Text()
		if band == "농도" {
			fmt.Printf("Review %d: %s - %s\n", i, band, title)
		}
	})
}

func main() {
	ExampleScrape()
}
