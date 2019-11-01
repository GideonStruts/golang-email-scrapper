package main

import (
	"fmt"
	"golang-email-scrapper/utils"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type GoogleResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"ke":  "https://www.google.co.ke/search?q=",
	"ug":  "https://www.google.co.ug/search?q=",
	"tz":  "https://www.google.co.tz/search?q=",
}

func buildGoogleUrl(searchTerm string, countryCode string, languageCode string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		return fmt.Sprintf("%s%s&num=100&hl=%s", googleBase, searchTerm, languageCode)
	} else {
		return fmt.Sprintf("%s%s&num=100&hl=%s", googleDomains["com"], searchTerm, languageCode)
	}
}

func googleRequest(searchURL string) (*http.Response, error) {

	baseClient := &http.Client{}

	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

	res, err := baseClient.Do(req)

	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func googleResultParser(response *http.Response) ([]GoogleResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	results := []GoogleResult{}
	sel := doc.Find("div.g")
	rank := 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h3.r")
		descTag := item.Find("span.st")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" {
			result := GoogleResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank += 1
		}
	}
	return results, err
}

func GoogleScrape(searchTerm string, countryCode string, languageCode string) ([]GoogleResult, error) {
	googleUrl := buildGoogleUrl(searchTerm, countryCode, languageCode)
	res, err := googleRequest(googleUrl)
	if err != nil {
		return nil, err
	}
	scrapes, err := googleResultParser(res)
	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func mainTwo() {

	results, _ := GoogleScrape("Top companies in East Africa", "TZ", "EN")

	for index, result := range results {
		fmt.Printf("%d : %s  : \n", index, result.ResultURL)

		links := utils.ScrapeOriginalUrlLinks(result.ResultURL)

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
}
