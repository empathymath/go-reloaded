package process

import (
	"strings"
)

// ApplyAtoAn μετατρέπει το "a" σε "an" όταν η επόμενη λέξη ξεκινά με φωνήεν ή h
func ApplyAtoAn(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for i := 0; i < len(words); i++ {
		// Αν η λέξη είναι "a" και υπάρχει επόμενη λέξη
		if words[i] == "a" && i+1 < len(words) {
			next := strings.ToLower(words[i+1])
			if strings.HasPrefix(next, "a") ||
				strings.HasPrefix(next, "e") ||
				strings.HasPrefix(next, "i") ||
				strings.HasPrefix(next, "o") ||
				strings.HasPrefix(next, "u") ||
				strings.HasPrefix(next, "h") {
				words[i] = "an"
			}
		}
		result = append(result, words[i])
	}

	return strings.Join(result, " ")
}
