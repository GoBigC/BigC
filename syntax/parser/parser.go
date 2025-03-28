package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"BigCooker/syntax/ast"
)

// Create a minimal Node implementation for your ast package
type MinimalNode struct{}

// ASTBuilder embeds the generated base visitor
type ASTBuilder struct {
	*BaseBigCVisitor
}

// NewASTBuilder creates a new ASTBuilder
func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{BaseBigCVisitor: &BaseBigCVisitor{}}
}

// Override just the VisitProgram method
func (v *ASTBuilder) VisitProgram(ctx *ProgramContext) interface{} {
	// Just return a minimal node implementation
	return &MinimalNode{}
}

func ParseFile(filename string) (ast.Node, error) {
	// 1. Stream input
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}

	// 2. Create lexer & parser instance 
	lexer := NewBigCLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewBigCParser(tokenStream)

	tree := parser.Program() // "program" is the grammar entrypoint

	// Use the constructor function
	builder := NewASTBuilder()
	astRoot := builder.VisitProgram(tree.(*ProgramContext))

	return astRoot.(ast.Node), nil
}