package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"BigCooker/syntax/ast"
)

func ParseFile(filename string) (*ast.Program, error) {
	// 1. Stream input
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}

	// 2. Create lexer & parser instance 
	lexer := NewBigCLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewBigCParser(tokenStream)

	tree := p.Program() // "program" is the grammar entrypoint

	// Use the constructor function
	builder := NewASTBuilder()
	astRoot := builder.VisitProgram(tree.(*ProgramContext))

	return astRoot.(*ast.Program), nil
}

type ASTBuilder struct {
	*BaseBigCVisitor
}

func NewASTBuilder() *ASTBuilder { // constructor
	return &ASTBuilder{BaseBigCVisitor: &BaseBigCVisitor{}}
}

func (v *ASTBuilder) VisitProgram(ctx *ProgramContext) interface{} {
	// set BaseNod properties
	program := &ast.Program {
		BaseNode: ast.BaseNode {
			Line: 	ctx.GetStart().GetLine(), 
			Column: ctx.GetStart().GetColumn(),
		},
	}

	// visit 
	n_declaration := 0
	for _, declCtx := range ctx.AllDeclaration() {
		decl := v.Visit(declCtx).(ast.Declaration)
		program.Declarations = append(program.Declarations, decl)
		n_declaration++
	}
	fmt.Printf("Found %d declarations\n", n_declaration)

	return program 
}