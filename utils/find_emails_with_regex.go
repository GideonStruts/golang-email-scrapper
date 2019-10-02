package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mcnijman/go-emailaddress"
)

func find_html_emails_with_regex() {

	url := "https://stackoverflow.com/questions/42407785/regex-extract-email-from-strings"

	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
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
		log.Fatal(err)
	}
	defer outFile.Close()

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// pulledData = "some data gidi@gmail.com alafu ingine "

	// https://github.com/mcnijman/go-emailaddress

	text := []byte(pulledData)
	validateHost := false

	email_s := emailaddress.Find(text, validateHost)

	for _, e := range email_s {
		fmt.Println(e)
	}

	// Create a regular expression to find emails
	// forOneEmailAddressRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	// re := regexp.MustCompile(`(?<name>[\w.]+)\@(?<domain>\w+\.\w+)(\.\w+)?`)
	// emails := re.FindAllString(pulledData, -1)
	// if emails == nil {
	// 	fmt.Println("No matches.")
	// } else {
	// 	for _, email := range emails {
	// 		fmt.Println(email)
	// 	}
	// }
}
