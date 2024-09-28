package kd_utils

import (
	"regexp"
)

func IsValidEmail(email *string) bool {

	// Define a regular expression for validating an email address.
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression.
	re := regexp.MustCompile(emailRegexPattern)

	// Match the email string against the regular expression.
	return re.MatchString(*email)
}
