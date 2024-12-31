package helper

import "regexp"

func CleanJSON(input string) string {
	whitespaceRegex := regexp.MustCompile(`\s+`)
	return whitespaceRegex.ReplaceAllString(input, "")
}
