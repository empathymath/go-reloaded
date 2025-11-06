package process

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	// ο αριθμός μετά το κόμμα είναι προαιρετικός πλέον
	cmdRe = regexp.MustCompile(`^\(\s*([A-Za-z]+)(?:\s*,\s*(\d+))?\s*\)$`)
	// θεωρούμε λέξη ό,τι περιέχει γράμμα ή αριθμό (Unicode digits/letters)
	wordRe  = regexp.MustCompile(`[\p{L}\p{N}]`)
	alphaRe = regexp.MustCompile(`\p{L}+`)
)

func isWord(token string) bool {
	return wordRe.MatchString(token)
}

func applyTransformToToken(tok, cmd string) string {
	lcmd := strings.ToLower(cmd)
	switch lcmd {
	case "up", "upper", "uppercase":
		return alphaRe.ReplaceAllStringFunc(tok, func(s string) string {
			return strings.ToUpper(s)
		})
	case "low", "lower", "lowercase":
		return alphaRe.ReplaceAllStringFunc(tok, func(s string) string {
			return strings.ToLower(s)
		})
	case "cap", "capitalize":
		return alphaRe.ReplaceAllStringFunc(tok, func(s string) string {
			if len(s) == 0 {
				return s
			}
			if len(s) == 1 {
				return strings.ToUpper(s)
			}
			return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
		})
	default:
		return tok
	}
}

// ApplyTextCommands εφαρμόζει εντολές τύπου (cmd[, N]) στις προηγούμενες N λέξεις.
// Μετράει και τα numeric tokens ως λέξεις, αλλά δεν αλλάζει τους αριθμούς — μόνο τα alphabetic runs.
func ApplyTextCommands(tokens []string) []string {
	out := make([]string, 0, len(tokens))
	appendedIndex := make(map[int]int)

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if m := cmdRe.FindStringSubmatch(t); m != nil {
			cmd := m[1]
			n := 1
			if m[2] != "" {
				if v, err := strconv.Atoi(m[2]); err == nil && v > 0 {
					n = v
				}
			}

			changed := 0
			j := i - 1
			for j >= 0 && changed < n {
				if isWord(tokens[j]) {
					tokens[j] = applyTransformToToken(tokens[j], cmd)
					if idx, ok := appendedIndex[j]; ok {
						out[idx] = tokens[j]
					}
					changed++
				}
				j--
			}
			// δεν εκπέμπουμε την εντολή
			continue
		}
		appendedIndex[i] = len(out)
		out = append(out, t)
	}
	return out
}
