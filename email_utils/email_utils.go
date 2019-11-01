package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mcnijman/go-emailaddress"
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

func ScrapeOriginalUrlLinks(url string) []string {
	var links []string
	links = ScrapeLinks(url)
	links = append(links, url)
	return links
}

func ScrapeURL(url string) []string {

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
