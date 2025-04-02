package semantic

import (
	"BigCooker/pkg/semantic/table"
	"BigCooker/pkg/syntax/ast"
)

// unlike other table structs, this struct is for the SemanticAnalyzer 
// itself, so it makes more sense to exist in this file
type SemanticAnalyzer struct {
	GlobalScope		*table.SymbolTable
	CurrentScope 	*table.SymbolTable
	Errors			[]string

	// this map holds type info for each expression 
	TypeMap	map[ast.Expression]table.TypeInfo
}

func New() *SemanticAnalyzer {
	// constructor logic

	return nil // placeholder return
	// a proper Go constructor for class Object should return 
	// the address to that object (&Object) with all 
	// fields properly initialized
}

func (a *SemanticAnalyzer) Analyze(program *ast.Program) []string {
	// analyzer logic

	return nil // placeholder return
}

// need a lot more functions, i just wrote function signature 
// for constructor and the main Analyzer there :) 