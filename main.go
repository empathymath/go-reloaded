package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/process"
)

func main() {
	// ✅ Έλεγχος ορισμάτων
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// ✅ Διαβάζουμε το αρχείο
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	// ✅ Tokenize το περιεχόμενο
	tokens := process.Tokenize(string(content))

	// ✅ Εφαρμόζουμε σταδιακά τους κανόνες
	tokens = process.ApplyHex(tokens)
	tokens = process.ApplyBin(tokens)
	tokens = process.ApplyTextCommands(tokens)

	// ✅ Επανασύνθεση πριν το punctuation stage
	text := strings.Join(tokens, " ")

	// ✅ Εφαρμόζουμε το punctuation fixer (δέχεται string, όχι []string)
	text = process.ApplyPunctuation(text)
	text = process.ApplyAtoAn(text)

	fmt.Println("DEBUG: before ApplyQuotes ->", text) // <-- add
	text = process.ApplyQuotes(text)
	fmt.Println("DEBUG: after  ApplyQuotes ->", text) // <-- add

	// ✅ Εγγραφή αποτελέσματος
	err = os.WriteFile(outputFile, []byte(text), 0o644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Println("✅ Text processed successfully!")
}
