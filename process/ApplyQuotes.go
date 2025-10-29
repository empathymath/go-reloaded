package process

import (
	"regexp"
	"strings"
)

// ApplyQuotes fixes single quotes and contractions
func ApplyQuotes(text string) string {
	// 1. Fix paired quotes (trim spaces inside)
	reQuotes := regexp.MustCompile(`'([^']*)'`)
	text = reQuotes.ReplaceAllStringFunc(text, func(match string) string {
		inner := strings.Trim(match, "' ")
		return "'" + inner + "'"
	})

	// 2. Fix contractions (remove spaces around internal apostrophes)
	// Example: don ' t → don't
	reContractions := regexp.MustCompile(`(\w+)\s+'\s+(\w+)`)
	text = reContractions.ReplaceAllString(text, `$1'$2`)

	return text
}
