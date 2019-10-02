package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeLinks(url string) []string {

	var links []string

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {

		linkTag := item
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()

		if len(link) > 5 && strings.Contains(link, "http") && strings.HasPrefix(link, "http") {
			fmt.Printf("Link: '%s' - '%s'\n", linkText, link)
			links = append(links, link)
		}
	})

	return links
}
