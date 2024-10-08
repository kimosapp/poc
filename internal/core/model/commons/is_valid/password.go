package is_valid

import (
	"unicode"
)

func IsValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 32 {
		return false
	}

	hasLower, hasUpper, hasDigit, hasSpecial := false, false, false, false

	for _, c := range password {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsDigit(c):
			hasDigit = true
		case isSpecialCharacter(c):
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}

func isSpecialCharacter(r rune) bool {
	specials := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	for _, s := range specials {
		if r == s {
			return true
		}
	}
	return false
}
