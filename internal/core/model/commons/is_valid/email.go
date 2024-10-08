package is_valid

import (
	"regexp"
	"strings"
)

func IsValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		return false
	}
	if !matched {
		return false
	}
	excludedPatterns := []string{
		`@123.123.123.123`,
		`@[123.123.123.123]`,
		`"^"`,
		`_@`,
	}
	for _, excluded := range excludedPatterns {
		if strings.Contains(email, excluded) {
			return false
		}
	}

	return true
}
