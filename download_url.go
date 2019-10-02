// download_url.go
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// Make request
	response, err := http.Get("https://stackoverflow.com/questions/42407785/regex-extract-email-from-strings")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

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
}
