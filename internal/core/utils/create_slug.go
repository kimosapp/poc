package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func CreateSlug(input string) string {
	slug := strings.ToLower(input)
	slug = strings.ReplaceAll(slug, " ", "-")
	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
	if err != nil {
		fmt.Println(err)
	}
	slug = reg.ReplaceAllString(slug, "")
	if len(slug) > 50 {
		slug = slug[:50]
	}

	return slug
}
