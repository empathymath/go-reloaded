package process

import (
	"regexp"
	"strings"
	"unicode"
)

// ApplyQuotes: δεν μετατρέπει αποστρόφους που είναι μέρος συστολής (don't) σε quotes.
// Αν το quote είναι πραγματικό ζευγάρι, trim-άρει τα άκρα του εσωτερικού περιεχομένου,
// προσθέτει κενό μετά σημείων στίξης πριν από άνοιγμα quote και καθαρίζει κενά πριν/μετά.
func ApplyQuotes(text string) string {
	// normalize common Unicode quote variants and invisible spaces
	repls := map[string]string{
		"‘": "'", "’": "'", "‚": "'", "‛": "'",
		"“": "\"", "”": "\"", "„": "\"",
		"\u00A0": " ",
		"\u200B": "",
		"\u200C": "",
		"\u200D": "",
		"\uFEFF": "",
	}
	for k, v := range repls {
		text = strings.ReplaceAll(text, k, v)
	}

	var b strings.Builder
	r := []rune(text)
	var lastRune rune // last rune written to builder

	writeRune := func(rr rune) {
		b.WriteRune(rr)
		lastRune = rr
	}
	writeString := func(s string) {
		if s == "" {
			return
		}
		b.WriteString(s)
		rs := []rune(s)
		lastRune = rs[len(rs)-1]
	}

	isPunctBefore := func(ch rune) bool {
		// treat common punctuation that should be followed by a space before an opening quote
		switch ch {
		case '.', ',', ':', ';', '!', '?', '(', '[', '{', '—', '–', '-', '/':
			return true
		}
		return false
	}

	for i := 0; i < len(r); {
		ch := r[i]

		// handle apostrophe inside words (possibly with spaces inserted by prior stages)
		if ch == '\'' || ch == '"' {
			quote := ch

			// detect contraction-like pattern: nearest non-space left and right are alnum -> apostrophe
			l := i - 1
			for l >= 0 && unicode.IsSpace(r[l]) {
				l--
			}
			rg := i + 1
			for rg < len(r) && unicode.IsSpace(r[rg]) {
				rg++
			}
			if l >= 0 && rg < len(r) && (unicode.IsLetter(r[l]) || unicode.IsDigit(r[l])) && (unicode.IsLetter(r[rg]) || unicode.IsDigit(r[rg])) {
				// treat as apostrophe/contraction, write the literal char and advance
				writeRune(ch)
				i++
				continue
			}

			// find matching closing quote
			j := i + 1
			for j < len(r) && r[j] != quote {
				j++
			}
			// no closing quote -> write as normal char
			if j >= len(r) {
				writeRune(ch)
				i++
				continue
			}

			// if previous written rune is punctuation and not a space, ensure a space before opening quote
			if lastRune != 0 && lastRune != ' ' && isPunctBefore(lastRune) {
				writeRune(' ')
			}

			// trim only edges inside quotes, keep interior spacing intact
			inner := strings.TrimSpace(string(r[i+1 : j]))
			writeRune(quote)
			writeString(inner)
			writeRune(quote)
			i = j + 1
			continue
		}

		writeRune(ch)
		i++
	}

	out := b.String()

	// remove space between closing quote and punctuation: 'something' . -> 'something'.
	reQuotePunct := regexp.MustCompile(`'([[:space:]]+)([.,;:?!])`)
	out = reQuotePunct.ReplaceAllString(out, `'$2`)
	reQuotePunctD := regexp.MustCompile(`"([[:space:]]+)([.,;:?!])`)
	out = reQuotePunctD.ReplaceAllString(out, `"$2`)

	// fallback: fix accidental spaces around apostrophes in contractions if any remain
	reContr := regexp.MustCompile(`([A-Za-z0-9])[[:space:]]*'[[:space:]]*([A-Za-z0-9])`)
	out = reContr.ReplaceAllString(out, `$1'$2`)

	return out
}
