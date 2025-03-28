package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"BigCooker/syntax/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bigcooker source_file")
		os.Exit(1)
	}

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	lexer := parser.NewBigCLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBigCParser(tokens)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	
	tree := p.Program() // start symbol is "program"

	// ----- Formatted tree -----
	// rawTree := tree.ToStringTree(p.RuleNames, p)
	// fmtTree := FormatParseTree(rawTree)
	// fmt.Println(fmtTree)

	// ----- Raw tree -----
	fmt.Print(tree.ToStringTree(p.RuleNames, p))
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