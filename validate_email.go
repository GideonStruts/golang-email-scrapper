package main

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func main() {

	email := "abc@abc12"

	if !validateEmail(email) {
		fmt.Println("Email address is invalid")
	} else {
		fmt.Println("Email address is VALID")
	}

	email = "abc@abc123.com"

	if !validateEmail(email) {
		fmt.Println("Email address is invalid")
	} else {
		fmt.Println("Email address is VALID")
	}
}
