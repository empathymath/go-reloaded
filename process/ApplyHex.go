package process

import (
	"fmt"
	"strconv"
)

// ApplyHex scans the tokens for "(hex)" and converts the previous token from hex to decimal.
func ApplyHex(tokens []string) []string {
	result := []string{}

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(hex)" {
			if len(result) == 0 {
				// Δεν υπάρχει προηγούμενη λέξη, αγνοούμε
				continue
			}

			// Προσπαθούμε να μετατρέψουμε την προηγούμενη λέξη από hex σε decimal
			hexStr := result[len(result)-1]
			val, err := strconv.ParseInt(hexStr, 16, 64)
			if err != nil {
				// Αν δεν είναι αριθμός, αγνοούμε
				fmt.Printf("Warning: cannot convert '%s' from hex\n", hexStr)
				continue
			}

			// Αντικαθιστούμε την προηγούμενη λέξη με το δεκαδικό αποτέλεσμα
			result[len(result)-1] = fmt.Sprintf("%d", val)
			continue
		}

		// Προσθέτουμε την τρέχουσα λέξη στο αποτέλεσμα
		result = append(result, tokens[i])
	}

	return result
}