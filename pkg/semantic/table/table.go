package table

import (
	"BigCooker/pkg/syntax/ast"
)

type SymbolTable struct {
	Parent 		*SymbolTable 
	Symbols 	map[string]Symbol
	ScopeType 	string 
}

type Symbol struct {
	Name 	string 
	Type	ast.Type	// "int", "char", float[]
	Kind 	string		// "variable", "function", "parameter"
	Line 	int
	Column 	int

	// functions only, nullable 
	Parameters []ast.Parameter

	IsInitialized bool 
	Used 		  bool
}

type TypeInfo struct {
	ExprType	ast.Type
	IsLValue	bool 	// make sure you understand LValue & RValue
}

// if a variable is used on the line outside of this range, 
// then throw a scope error
type ScopeInfo struct {
	ValidFirstLine 	int 
	ValidLastLine	int
}