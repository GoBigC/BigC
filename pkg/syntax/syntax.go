package syntax

import (
	"bytes"
	"os"

	"BigCooker/pkg/syntax/ast"
	"BigCooker/pkg/syntax/parser"
)

func ParseFile(filename string) (*ast.Program, error) {
	return parser.ParseFile(filename)
}

func GetCST(filename string) string {
	return getCST(filename)
}

func ProcessFile(filename string) error {
	outputDir := "artifact"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err // Uncomment this to properly handle directory creation errors
	}

	cst := getCST(filename)
	cstFormatted := FormatParseTree(cst)
	if err := os.WriteFile(outputDir+"/cst.txt", []byte(cstFormatted), 0644); err != nil {
		return err // Uncomment this to properly handle file write errors
	}

	program, err := parser.ParseFile(filename)
	if err != nil {
		// If there's a parsing error, write the error to a file and return
		if err := os.WriteFile(outputDir+"/parse_error.txt", []byte(err.Error()), 0644); err != nil {
			return err
		}
		return err // Return here to prevent using the nil program
	}

	// Only proceed with AST operations if parsing was successful
	var astBuffer bytes.Buffer
	astBuffer.WriteString("========== AST ==========\n")
	PrintAST(program, "", &astBuffer)
	if err := os.WriteFile(outputDir+"/ast.txt", astBuffer.Bytes(), 0644); err != nil {
		return err // Uncomment this to properly handle file write errors
	}

	if err := DrawASTTree(program, outputDir+"/ast_tree.txt"); err != nil {
		return err // Uncomment this to properly handle tree drawing errors
	}

	if err := DrawCSTTree(filename, outputDir+"/cst_tree.txt"); err != nil {
		return err // Uncomment this to properly handle tree drawing errors
	}

	return nil
}
