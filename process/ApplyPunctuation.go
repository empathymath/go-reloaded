package process

import (
	"regexp"
	"strings"
)

// ApplyPunctuation διορθώνει αποστάσεις και μορφοποίηση στίξης.
func ApplyPunctuation(text string) string {
	// 1️⃣ Ενιαία σύνολα στίξης που πρέπει να μείνουν "ενωμένα"
	text = strings.ReplaceAll(text, " . . .", "...")
	text = strings.ReplaceAll(text, ". . .", "...")
	text = strings.ReplaceAll(text, " ! ?", "!?")
	text = strings.ReplaceAll(text, " ? !", "?!")

	// 2️⃣ Αφαιρεί τα κενά πριν από σημεία στίξης
	re := regexp.MustCompile(`\s+([,.!?;:])`)
	text = re.ReplaceAllString(text, "$1")

	// 3️⃣ Εξασφαλίζει ότι υπάρχει ένα (και μόνο ένα) space μετά από σημείο στίξης
	re = regexp.MustCompile(`([,.!?;:])([^\s,.!?;:])`)
	text = re.ReplaceAllString(text, "$1 $2")

	// 4️⃣ Καθαρίζει διπλά spaces
	re = regexp.MustCompile(`\s{2,}`)
	text = re.ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}
