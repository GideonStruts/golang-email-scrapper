package utils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func downloadURL(url string) {

	// Make request
	response, err := http.Get(url)
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
