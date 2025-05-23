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

func ProcessFile(filename string)(*ast.Program, error) {
    outputDir := "artifact"

    if err := os.MkdirAll(outputDir, 0755); err != nil {
        return nil, err
    }
    
    cst := getCST(filename)
    cstFormatted := FormatParseTree(cst)
    if err := os.WriteFile(outputDir+"/cst.txt", []byte(cstFormatted), 0644); err != nil {
        return nil, err
    }
    
    program, err := parser.ParseFile(filename)
    if err != nil {
        return nil, err
    }
    
    var astBuffer bytes.Buffer
    astBuffer.WriteString("========== AST ==========\n")
    PrintAST(program, "", &astBuffer)
    if err := os.WriteFile(outputDir+"/ast.txt", astBuffer.Bytes(), 0644); err != nil {
        return nil, err
    }
    
    if err := DrawASTTree(program, outputDir+"/ast_tree.txt"); err != nil {
        return nil, err
    }
    
    if err := DrawCSTTree(filename, outputDir+"/cst_tree.txt"); err != nil {
        return nil, err
    }
    
    return program, nil
}