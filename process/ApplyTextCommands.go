package process

import (
	"fmt"
	"strings"
	"unicode"
)

// ApplyTextCommands handles (up), (low), (cap), (up, N), (low, N), (cap, N)
func ApplyTextCommands(tokens []string) []string {
	result := []string{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Ελέγχουμε αν είναι εντολή
		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {
			cmd := strings.TrimSuffix(strings.TrimPrefix(token, "("), ")")
			parts := strings.Split(cmd, ",")
			action := strings.TrimSpace(parts[0])
			count := 1

			// Αν υπάρχει αριθμός (π.χ. (up, 3))
			if len(parts) == 2 {
				fmt.Sscanf(parts[1], "%d", &count)
			}

			// Αν δεν υπάρχουν προηγούμενες λέξεις, συνεχίζουμε
			if len(result) == 0 {
				continue
			}

			// Εφαρμόζουμε τη μετατροπή στις προηγούμενες λέξεις
			for j := len(result) - 1; j >= 0 && count > 0; j-- {
				switch action {
				case "up":
					result[j] = strings.ToUpper(result[j])
				case "low":
					result[j] = strings.ToLower(result[j])
				case "cap":
					result[j] = capitalize(result[j])
				default:
					fmt.Printf("Unknown command: %s\n", action)
				}
				count--
			}

			continue
		}

		result = append(result, token)
	}

	return result
}

// capitalize μετατρέπει την πρώτη ρουτίνα ενός string σε κεφαλαίο γράμμα
func capitalize(word string) string {
	if word == "" {
		return ""
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}
