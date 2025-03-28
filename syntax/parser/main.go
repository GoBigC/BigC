package parser 

import (
	"github.com/antlr4-go/antlr/v4"
	"BigCooker/syntax/ast"
)
type ASTBuilder struct {
	BaseBigCVisitor
}

func ParseFile(filename string) (ast.Node, error) {
	// 1. Stream input
	input, err := antlr.NewFileStream(filename)
	if (err != nil) { 
		return nil, err
	}

	// 2. Create lexer & parser instance 
	lexer := NewBigCLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewBigCParser(tokenStream)

	tree := parser.Program() // "program" is the grammar entrypoint

	builder := &ASTBuilder{}
	astRoot := builder.Visit(tree)

	return astRoot.(ast.Node), nil 
}