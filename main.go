package main

import (
	"fmt"
	"os"
	// "go-reloaded/internal/pipeline"
)

func main() {
	// Validate command-line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}
	
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	
	// διαβαζει το αρχειο
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
		os.Exit(1)
	}
	
	// // στο μελλον θα γινεται επεξεργασία
	result := string(content)
	
	// εγγραφή στο αρχείο αποτελέσματος
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outputFile, err)
		os.Exit(1)
	}
	
	fmt.Println("Text processed successfully!")
}