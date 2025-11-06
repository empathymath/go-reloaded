package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"go-reloaded/process"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	// Χωρίζουμε το input σε γραμμές
	lines := strings.Split(string(content), "\n")
	processedLines := make([]string, 0, len(lines))

	// Επεξεργαζόμαστε κάθε γραμμή ξεχωριστά
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			processedLines = append(processedLines, line) // διατήρηση κενών γραμμών
			continue
		}

		tokens := process.Tokenize(line)
		tokens = process.ApplyHex(tokens)
		tokens = process.ApplyBin(tokens)
		tokens = process.ApplyTextCommands(tokens)

		text := strings.Join(tokens, " ")
		text = process.ApplyPunctuation(text)
		text = process.ApplyAtoAn(text)
		text = process.ApplyQuotes(text)

		// cleanup: remove space after opening quotes when followed by a letter
		reSingle := regexp.MustCompile(`'\s+([A-Za-z])`)
		text = reSingle.ReplaceAllString(text, `'$1`)
		reDouble := regexp.MustCompile(`"\s+([A-Za-z])`)
		text = reDouble.ReplaceAllString(text, `"$1`)

		processedLines = append(processedLines, text)
	}

	// Επανασύνθεση με τις αλλαγές γραμμής
	output := strings.Join(processedLines, "\n")

	// Προσθήκη newline στο τέλος αν υπήρχε στο input
	if len(content) > 0 && content[len(content)-1] == '\n' {
		output += "\n"
	}

	err = os.WriteFile(outputFile, []byte(output), 0o644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Println("✅ Text processed successfully!")
}
