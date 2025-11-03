package process

import (
	"strings"
	"unicode"
)

// ApplyAtoAn μετατρέπει το "a" σε "an" όταν η επόμενη λέξη ξεκινά με φωνήεν ή h
func ApplyAtoAn(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for i := 0; i < len(words); i++ {
		// Αν η λέξη είναι "a" (case-insensitive) και υπάρχει επόμενη λέξη
		if strings.EqualFold(words[i], "a") && i+1 < len(words) {
			next := strings.ToLower(words[i+1])
			if strings.HasPrefix(next, "a") ||
				strings.HasPrefix(next, "e") ||
				strings.HasPrefix(next, "i") ||
				strings.HasPrefix(next, "o") ||
				strings.HasPrefix(next, "u") ||
				strings.HasPrefix(next, "h") {
				// Preserve original capitalization of the "a"
				if len(words[i]) > 0 && unicode.IsUpper(rune(words[i][0])) {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
		result = append(result, words[i])
	}

	return strings.Join(result, " ")
}
