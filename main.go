package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/process"
)

func main() {
	// Validate command-line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Διαβάζουμε το αρχείο
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	// Καλούμε το Tokenizer για να πάρουμε τα tokens
	tokens := process.Tokenize(string(content))
	// 2Εφαρμόζουμε τον κανόνα (hex)
	tokens = process.ApplyHex(tokens)
	tokens = process.ApplyBin(tokens)

	// Επανασυνθέτουμε τα tokens σε ένα κείμενο με απλό διάστημα μεταξύ
	processedText := strings.Join(tokens, " ")

	// Εγγραφή στο αρχείο αποτελέσματος
	err = os.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Println("Text tokenized successfully!")
}
