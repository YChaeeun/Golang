package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape1() {
	// Request the HTML page.
	res, err := http.Get("https://comic.naver.com/webtoon/weekday.nhn")
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
	doc.Find(".list_area").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		weekday := s.Find(".mon").Text()
		fmt.Printf("%s\n", weekday)
		title := s.Find(".mon + ul li .thumb + a").Text()
		fmt.Printf("%s\n", title)
	})

	doc.Find(".asideBox").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		boxtitle := s.Find(".asideBoxTitle + div ul + ol li a").Text()
		fmt.Printf("%s\n", boxtitle)
		//title := s.Find(".mon + ul li .thumb + a").Text()
		//fmt.Printf("%s\n", title)
	})
}

func main() {
	ExampleScrape1()
}
