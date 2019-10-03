package main

import (
	"fmt"
	"log"
	"strings"
	"web-scraping/colly"
	"web-scraping/utils"
)

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
