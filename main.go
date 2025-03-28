package main

import (
	"fmt"
	"os"
	
	// "BigCooker/syntax"
	"BigCooker/syntax/parser"
	// "BigCooker/syntax/ast"

	// "BigCooker/semantics"
	// "BigCooker/codegen"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Usage: bigcooker <source-file-name>")
		os.exit(1)
	}
	sourceFile := os.Args[1]

	// 1. Syntax 
	astRoot, err := parser.ParseFile(sourceFile)
	if (err != nil) {
		fmt.Printf("Error parsing file: %v\n", err)
		os.exit(1)
	}

	// 2. Semantics 

	// 3. Code generation 

	fmt.Println("Compilation completed.")
}