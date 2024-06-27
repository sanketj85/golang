package controllers

import (
	"regexp"
)

func isValidName(name string) bool {
	// Regular expression to match alphabetic characters only
	regex := regexp.MustCompile("^[a-zA-Z ]+$")
	return regex.MatchString(name)
}

func isValidPhoneNumber(phoneNumber string) bool {
	// Regular expression to match 10 numeric digits
	regex := regexp.MustCompile(`^\d{10}$`)
	return regex.MatchString(phoneNumber)
}

func isValidEmail(email string) bool {
	// Regular expression to match a valid email address format
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return regex.MatchString(email)
}
