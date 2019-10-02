package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"web-scraping/colly"
	"web-scraping/utils"

	"github.com/mcnijman/go-emailaddress"
)

func main() {

	colly.TestColly()

	url := "https://www.theeastafrican.co.ke/"

	links := scrapeLinks(url)

	var emails []string

	// counter := 0

	for index, link := range links {

		fmt.Printf("%d. %s\n", index, link)
		linkEmails := scrapeURL(link)

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

func scrapeLinks(url string) []string {
	var links []string
	links = utils.ScrapeLinks(url)
	links = append(links, url)
	return links
}

func scrapeURL(url string) []string {

	var emails []string

	if len(url) < 5 || !strings.Contains(url, "http") || url == "http://www.ntv.co.ke" {
		return emails
	}

	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error processing URL  >>> ", err)
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading HTTP body. ", err)
	}

	pulledData := string(body)
	pulledData = strings.Replace(pulledData, "<", "", -1)
	pulledData = strings.Replace(pulledData, ">", "", -1)
	pulledData = strings.Replace(pulledData, "\"", "", -1)
	pulledData = strings.Replace(pulledData, "/", "", -1)

	// fmt.Println("Response Body >>>>  ", pulledData)

	// Create output file
	outFile, err := os.Create("output.txt")
	if err != nil {
		log.Println(err)
	}
	defer outFile.Close()

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Println(err)
	}

	// pulledData = "some data gidi@gmail.com alafu ingine "

	// https://github.com/mcnijman/go-emailaddress

	text := []byte(pulledData)
	validateHost := false

	email_s := emailaddress.Find(text, validateHost)

	for _, email := range email_s {
		emails = append(emails, email.String())
	}

	return emails
}
