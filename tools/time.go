package tools

import (
	"time"
)

// CheckTimeFormat checks if the given time layout is valid. It returns true
// if the given layout is valid and false otherwise.
func CheckTimeFormat(layout string) bool {
	if _, err := time.Parse(layout, "2022-07-25 14:30:00"); err != nil {
		return true
	}
	return false
}
