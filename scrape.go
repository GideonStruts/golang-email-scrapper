package main

import (
	"fmt"
	"golang-email-scrapper/colly"
	"golang-email-scrapper/utils"
	"log"
	"strings"
)

// func main() {
// 	c := colly.NewCollector()

// 	// Find and visit all links
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		e.Request.Visit(e.Attr("href"))
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	c.Visit("http://go-colly.org/")
// }

func mainOne() {

	colly.TestColly()

	url := "https://www.theeastafrican.co.ke/"

	links := utils.ScrapeOriginalUrlLinks(url)

	var emails []string

	// counter := 0

	for index, link := range links {

		fmt.Printf("%d. %s\n", index, link)
		linkEmails := utils.ScrapeURL(link)

		for _, email := range linkEmails {

			if !strings.HasSuffix(email, ".png") && len(email) < 60 {
				// fmt.Printf("%d. Appending email >>>  %s\n", index, email)
				emails = append(emails, email)
			}
		}

	}

	log.Println("\n\n=======================================================================\n\n")
	log.Println("Printing all emails  >>>>   ", len(emails))
	log.Println("\n\n=======================================================================\n\n")

	for index, email := range emails {
		fmt.Printf("%d. %s\n", index+1, email)
	}

	log.Println("\n\n=======================================================================\n\n")

}
