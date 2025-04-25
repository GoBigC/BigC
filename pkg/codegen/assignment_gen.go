package codegen

import (
	"BigCooker/pkg/semantic/table"
	"BigCooker/pkg/syntax/ast"
	"fmt"
)

type AssignmentGenerator struct {
	CodeGen *CodeGenerator
	SymTab  *table.SymbolTable
}

func NewAssignmentGenerator(cg *CodeGenerator, symtab *table.SymbolTable) *AssignmentGenerator {
	return &AssignmentGenerator{
		CodeGen: cg,
		SymTab:  symtab,
	}
}

func (ag *AssignmentGenerator) GenerateVarDeclaration() {

	for _, symbol := range ag.SymTab.Symbols {
		if symbol.ReturnType != nil {
			continue // Skip function symbols
		}

		switch symbolType := symbol.Type.(type) {
		case *ast.PrimitiveType:
			if symbol.Value != nil {
				ag.CodeGen.insertData(symbol.Name, ".word", symbol.Value)
			} else {
				ag.CodeGen.insertData(symbol.Name, ".space", 8)
			}
		case *ast.ArrayType:
			ag.CodeGen.insertData(symbol.Name, ".space", 8*symbol.ArraySize)
		default:
			panic(fmt.Sprintf("Unsupported type for variable declaration: %T", symbolType))
		}
	}
}
