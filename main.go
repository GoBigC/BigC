package main

import (
	"BigCooker/pkg/error_formatter"
	"BigCooker/pkg/semantic"
	"BigCooker/pkg/syntax"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please supply source file")
		os.Exit(1)
	}
	filename := os.Args[1]

	sourceBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	sourceText := string(sourceBytes)
	sourceLines := strings.Split(sourceText, "\n")

	program, err := syntax.ProcessFile(filename)
	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully processed file and generated artifacts.")
	fmt.Println("Syntax analysis completed successfully.")

	// Initialize the semantic analyzer
	semanticAnalyzer := semantic.NewSemanticAnalyzer()

	// Perform semantic analysis
	if errs := semanticAnalyzer.Analyze(program); len(errs) > 0 {
		errorFormatter := error_formatter.NewSemanticErrorFormatter(sourceLines)
		formattedErrors := errorFormatter.FormatErrors(errs)

		for _, formattedErr := range formattedErrors {
			fmt.Println(formattedErr)
			fmt.Println()
		}
		os.Exit(1)
	}
	fmt.Println("Semantic analysis completed successfully.")
}
