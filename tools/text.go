package tools

import "regexp"

// StringContainsNumber checks if the given text contains any numeric digits.
// It returns true if the text contains at least one digit, otherwise false.
func StringContainsNumber(text string) bool {
	re := regexp.MustCompile(`\d+`)
	if len(text) > 0 && re.MatchString(text) {
		return true
	}
	return false
}
