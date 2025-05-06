package main

import (
	"BigCooker/pkg/codegen"
	"BigCooker/pkg/error_formatter"
	"BigCooker/pkg/semantic"
	"BigCooker/pkg/syntax"
	"fmt"
	"os"
	"strings"
)

func main() {
	// ----- PREPARING IN/OUT FILES -----
	if len(os.Args) < 2 {
		fmt.Println("Please supply source file")
		os.Exit(1)
	}
	filename := os.Args[1]

	sourceBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	sourceText := string(sourceBytes)
	sourceLines := strings.Split(sourceText, "\n")

	// ----- PHASE 1: SYNTAX ANALYSIS -----
	program, err := syntax.ProcessFile(filename)
	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully processed file and generated artifacts.")
	fmt.Println("Syntax analysis completed successfully.")

	// ----- PHASE 2: SEMANTIC ANALYSIS -----
	semanticAnalyzer := semantic.NewSemanticAnalyzer()

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

	// ----- PHASE 3: CODE GENERATION -----
	// symbolTables := semanticAnalyzer.SymTable
	// ast := program
	outFile := "asm.asm"
	codeGenerator := codegen.NewCodeGenerator(program, semanticAnalyzer.SymTable)

	if err := codeGenerator.GenerateProgram(outFile); err != nil {
		fmt.Printf("Error generating code: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Code generation completed successfully.")
}
