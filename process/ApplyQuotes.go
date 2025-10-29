package process

import (
	"regexp"
	"strings"
)

// ApplyQuotes fixes single quotes around words or phrases
func ApplyQuotes(text string) string {
	// Regex to match anything inside single quotes (including spaces)
	re := regexp.MustCompile(`'([^']*)'`)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		// Remove leading/trailing spaces inside the quotes
		inner := strings.Trim(match, "' ") // remove ' and spaces
		return "'" + inner + "'"           // put quotes back tight
	})
}
