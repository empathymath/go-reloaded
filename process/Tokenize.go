package process
import (
	"regexp"
	"strings"
)

// Tokenize παίρνει ένα κείμενο και το σπάει σε tokens
func Tokenize(text string) []string {
	// Καθαρίζουμε τα περιττά whitespaces
	text = strings.TrimSpace(text)

	// Κανονική έκφραση για tokens:
	// Λέξεις, αριθμούς, εντολές σε παρένθεση, στίξη, quotes
	re := regexp.MustCompile(`\w+|\(\w+(?:,\s*\d+)?\)|[.,!?;:'’]+`)

	tokens := re.FindAllString(text, -1)
	return tokens
}