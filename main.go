package main

import (
	"database/sql"
	"fmt"
	"golang-email-scrapper/utils"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	var emails []string
	counter := 1
	db := CreateMySqlCon()

	c.OnRequest(func(r *colly.Request) {
		link := r.URL.String()
		fmt.Println(counter, ": Visiting", link)

		linkEmails := utils.ScrapeURL(link)

		for _, email := range linkEmails {

			if !strings.HasSuffix(email, ".png") && len(email) < 60 {
				// fmt.Printf("%d. Appending email >>>  %s\n", index, email)
				emails = append(emails, email)
			}
		}

		if counter >= 10 {
			printEmails(emails, db)
		}
		counter++
	})

	c.Visit("http://go-colly.org/")

}

func printEmails(emails []string, db *sql.DB) {

	log.Println("\n\n=======================================================================\n\n")
	log.Println("Printing all emails  >>>>   ", len(emails))
	log.Println("\n\n=======================================================================\n\n")

	for index, email := range emails {
		fmt.Printf("%d. %s\n", index+1, email)

		// perform a db.Query insert
		insert, err := db.Query("INSERT INTO emails(email_address, scrape_url) VALUES ('" + email + "', '" + email + "')")

		// if there is an error inserting, handle it
		if err != nil {
			// panic(err.Error())
			log.Println("Error saving email address to db : ", err)
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
	}

	log.Println("\n\n=======================================================================\n\n")
}

/*Create mysql connection*/
func CreateMySqlCon() *sql.DB {
	db, err := sql.Open("mysql", "gidi:smallz@tcp(localhost:3306)/scraped_emails")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
