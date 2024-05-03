package helper

import (
	"fmt"
	"regexp"
)

func IsValidEmail(email string) bool {
	// Regular expression for validating email addresses without allowing special characters in the local part
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func RecoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
