package process

import (
	"fmt"
	"strconv"
)

// ApplyHex: Σαρώνει τα tokens για το "(hex)" και μετατρέπει το προηγούμενο token από δεκαεξαδικό σε δεκαδικό.
func ApplyHex(tokens []string) []string {
	result := []string{}

	for i := 0; i < len(tokens); i++ {
		// Αν το token είναι "(hex)", μετατρέπουμε το προηγούμενο token από δεκαεξαδικό σε δεκαδικό
		if tokens[i] == "(hex)" {
			if len(result) == 0 {
				// Δεν υπάρχει προηγούμενη λέξη, αγνοείται
				continue
			}

			// Μετατροπή της προηγούμενης λέξης από hex σε decimal
			hexStr := result[len(result)-1]
			val, err := strconv.ParseInt(hexStr, 16, 64)
			if err != nil {
				// Αν δεν είναι αριθμός, αγνοείται
				fmt.Printf("Προειδοποίηση: δεν ήταν δυνατή η μετατροπή του '%s' από δεκαεξαδικό\n", hexStr)
				continue
			}

			// Αντικαθιστούμε την προηγούμενη λέξη με το δεκαδικό αποτέλεσμα
			result[len(result)-1] = fmt.Sprintf("%d", val)
			continue
		}

		// Προσθέτουμε το τρέχον token στο αποτέλεσμα
		result = append(result, tokens[i])
	}

	// Επιστρέφουμε το νέο slice με τις μετατροπές
	return result
}
