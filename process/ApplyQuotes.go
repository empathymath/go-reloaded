package process

import (
	"regexp"
	"strings"
	"unicode"
)

// ApplyQuotes:
// - Δεν πειράζει καθόλου συστολές (don't, I'm, you're)
// - Διορθώνει quotes του project (' ... ')
// - Κάνει trim τα κενά ΜΕΣΑ στα quotes
// - Ενώνει σωστά quotes με σημεία στίξης
// - Δεν κρασάρει ποτέ (safe slice handling)
func ApplyQuotes(text string) string {
	runes := []rune(text)
	n := len(runes)

	var out strings.Builder

	inQuote := false         // true όταν βρισκόμαστε μέσα σε quotes
	quoteChar := rune(0)     // ' ή "
	quoteStartIndex := 0     // το index στο out όπου ξεκινάει το quote

	for i := 0; i < n; i++ {
		ch := runes[i]

		// 1. Αν πρόκειται για συστολή (contraction) όπως don't, I'm, you're
		// τότε η απόστροφος βρίσκεται ανάμεσα σε δύο γράμματα
		// και ΔΕΝ πρέπει να την πειράξουμε.
		if ch == '\'' {
			if i > 0 && i+1 < n &&
				unicode.IsLetter(runes[i-1]) &&
				unicode.IsLetter(runes[i+1]) {

				out.WriteRune(ch)
				continue
			}
		}

		// 2. Αν το χαρακτήρας είναι quote (' ή ")
		if ch == '\'' || ch == '"' {

			// --- ΑΝΟΙΓΜΑ QUOTE ---
			if !inQuote {
				inQuote = true
				quoteChar = ch
				quoteStartIndex = out.Len() // αποθηκεύουμε πού ξεκίνησε
				out.WriteRune(ch)
				continue
			}

			// --- ΚΛΕΙΣΙΜΟ QUOTE ---
			if inQuote && ch == quoteChar {
				inQuote = false

				// Παίρνουμε όλο το ήδη γραμμένο κείμενο
				full := out.String()

				// Το μέρος πριν το opening quote (συμπεριλαμβάνει το quote)
				before := full[:quoteStartIndex+1]

				// Το εσωτερικό των quotes
				inner := strings.TrimSpace(full[quoteStartIndex+1:])

				// Χτίζουμε από την αρχή το νέο out
				out.Reset()
				out.WriteString(before)
				out.WriteString(inner)
				out.WriteRune(ch)

				continue
			}
		}

		// Κανονική εγγραφή χαρακτήρα
		out.WriteRune(ch)
	}

	result := out.String()

	// 3. Αφαιρούμε επιπλέον κενά μετά το άνοιγμα quote:
	// '   word → 'word
	re1 := regexp.MustCompile(`'\s+([A-Za-z])`)
	result = re1.ReplaceAllString(result, `'$1`)

	re2 := regexp.MustCompile(`"\s+([A-Za-z])`)
	result = re2.ReplaceAllString(result, `"$1`)

	// 4. Αφαιρούμε κενά πριν από σημεία στίξης:
	// 'word ' ! → 'word!
	re3 := regexp.MustCompile(`'(\s+)([.,!?;:])`)
	result = re3.ReplaceAllString(result, `'$2`)

	re4 := regexp.MustCompile(`"(\s+)([.,!?;:])`)
	result = re4.ReplaceAllString(result, `"$2`)

	return result
}
