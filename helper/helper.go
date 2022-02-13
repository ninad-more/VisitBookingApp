package helper

import (	
	"regexp"
)

func IsValidInputs(firstName string, lastName string, email string) (bool,bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isValidEmail := emailRegex.MatchString(email)
	
	return isvalidName, isValidEmail
}