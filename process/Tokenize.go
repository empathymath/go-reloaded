package process

import (
	"regexp"
	"strings"
)

// Tokenize σπάει σωστά το κείμενο σε tokens χωρίς να χαλάει contractions
func Tokenize(text string) []string {
	text = strings.TrimSpace(text)

	// 1. regex για tokens:
	// λέξεις, αριθμοί, εντολές σε παρένθεση, punctuation, quotes
	re := regexp.MustCompile(`\w+|\(\w+(?:,\s*\d+)?\)|[.,!?;:'’]+`)
	tokens := re.FindAllString(text, -1)

	// 2. merge contractions: "don" "'" "t" → "don't"
	merged := []string{}
	i := 0
	for i < len(tokens) {

		if i+2 < len(tokens) &&
			isWordToken(tokens[i]) &&
			isAposToken(tokens[i+1]) &&
			isWordToken(tokens[i+2]) {

			merged = append(merged, tokens[i]+"'" + tokens[i+2])
			i += 3
			continue
		}

		merged = append(merged, tokens[i])
		i++
	}

	return merged
}

// --- helpers ειδικά για το Tokenize ---
// (δεν συγκρούονται πλέον με isWord() από ApplyTextCommands.go)

func isWordToken(s string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(s)
}

func isAposToken(s string) bool {
	return s == "'" || s == "’"
}
