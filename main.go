package main

import (
	"fmt"
	"os"
	"strings"

	"BigCooker/syntax/parser"
	"BigCooker/syntax/ast"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please supply source file")
		os.Exit(1)
	}
	program, err := parser.ParseFile(os.Args[1])
	if (err != nil) {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("========== AST ==========")
	PrintAST(program, "")
}

func PrintAST(node ast.Node, indent string){
	switch n:= node.(type) {
	case *ast.Program:
		fmt.Printf("%sProgram (Line %d, Col %d) with %d declarations\n",
			indent, n.Line, n.Column, len(n.Declarations))
		for i, decl := range n.Declarations {
			fmt.Printf("%sDeclaration %d:\n", indent+" ", i+1)
			PrintAST(decl, indent+" ")
		}
	default: 
		fmt.Printf("%sUnknown node type: %T\n", indent, node)
	}
}

func FormatParseTree(rawTree string) string {
	rawTree = strings.TrimSpace(rawTree)
	var fmtTree strings.Builder 
	n_indent := 0 
	inParen := false 

	for i := 0; i < len(rawTree); i++ {
		char := rawTree[i]

		switch (char) { 
		case '(': 
			fmtTree.WriteByte(char)
			n_indent++
			inParen = true 

			fmtTree.WriteString("\n")
			fmtTree.WriteString(strings.Repeat("\t", n_indent))

		case ')': 
			n_indent-- 
			if (inParen) {
				fmtTree.WriteString("\n")
				fmtTree.WriteString(strings.Repeat("\t", n_indent))
			}
			fmtTree.WriteByte(char)
			inParen = false 

		case ' ': 
			if (i > 0 && rawTree[i-1]=='('){
				continue
			}

			if (!inParen){
				fmtTree.WriteString("\n")
				fmtTree.WriteString(strings.Repeat("\t", n_indent))
			} else {
				fmtTree.WriteByte(char)
			}
		
		default: 
			fmtTree.WriteByte(char)
		}
	}

	return fmtTree.String()
}