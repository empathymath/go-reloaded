package process

import (
	"fmt"
	"strconv"
)

// ApplyBin scans the tokens for "(bin)" and converts the previous token from binary to decimal.
func ApplyBin(tokens []string) []string {
	result := []string{}

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(bin)" {
			if len(result) == 0 {
				continue
			}

			binStr := result[len(result)-1]
			val, err := strconv.ParseInt(binStr, 2, 64)
			if err != nil {
				fmt.Printf("Warning: cannot convert '%s' from binary\n", binStr)
				continue
			}

			result[len(result)-1] = fmt.Sprintf("%d", val)
			continue
		}

		result = append(result, tokens[i])
	}

	return result
}
