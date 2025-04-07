package main

import (
	"fmt"
	"os"

	"BigCooker/pkg/syntax"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please supply source file")
		os.Exit(1)
	}

	filename := os.Args[1]

	err := syntax.ProcessFile(filename)
	if err != nil {
		fmt.Printf("Syntax error(s): %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully processed file and generated artifacts.")
}
