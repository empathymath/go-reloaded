package process

import (
	"fmt"
	"strconv"
)

// ApplyBin: Σαρώνει τα tokens για το "(bin)" και μετατρέπει το προηγούμενο token από δυαδικό σε δεκαδικό.
func ApplyBin(tokens []string) []string {
	result := []string{}

	for i := 0; i < len(tokens); i++ {
		// Αν το token είναι "(bin)", μετατρέπουμε το προηγούμενο token από δυαδικό σε δεκαδικό
		if tokens[i] == "(bin)" {
			if len(result) == 0 {
				continue
			}

			binStr := result[len(result)-1]
			val, err := strconv.ParseInt(binStr, 2, 64)
			if err != nil {
				fmt.Printf("Προειδοποίηση: δεν ήταν δυνατή η μετατροπή του '%s' από δυαδικό\n", binStr)
				continue
			}

			result[len(result)-1] = fmt.Sprintf("%d", val)
			continue
		}

		// Προσθέτουμε το token στο αποτέλεσμα
		result = append(result, tokens[i])
	}

	// Επιστρέφουμε το νέο slice με τις μετατροπές
	return result
}
