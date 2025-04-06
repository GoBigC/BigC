package main

import (
	"fmt"
	"os"

	"BigCooker/pkg/semantic"
	"BigCooker/pkg/syntax"
)

func main() {
    // if len(os.Args) < 2 {
    //     fmt.Println("Please supply source file")
    //     os.Exit(1)
    // }
    
    // filename := os.Args[1]
    filename := "test/smol_sample.uia"
    
    program, err := syntax.ProcessFile(filename)
    if err != nil {
        fmt.Printf("Error processing file: %v\n", err)
        os.Exit(1)
    }

    
    fmt.Println("Successfully processed file and generated artifacts.")

    // Initialize the semantic analyzer
    semanticAnalyzer := semantic.NewSemanticAnalyzer()
    // Perform semantic analysis
    if err := semanticAnalyzer.Analyze(program); len(err) > 0 {
        fmt.Printf("Semantic analysis error: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Semantic analysis completed successfully.")

}