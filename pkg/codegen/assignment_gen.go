package codegen

import (
	"BigCooker/pkg/syntax/ast"
	"fmt"
)

type AssignmentGenerator struct {
	CodeGen *CodeGenerator
}

func NewAssignmentGenerator(cg *CodeGenerator) *AssignmentGenerator {
	return &AssignmentGenerator{
		CodeGen: cg,
	}
}

func (ag *AssignmentGenerator) GenerateVarDeclaration() {

	for _, symbol := range ag.CodeGen.SymTable.Symbols {
		if symbol.ReturnType != nil || symbol.Parameters != nil {
			continue // Skip function symbols
		}

		switch symbolType := symbol.Type.(type) {
		case *ast.PrimitiveType:
			if symbol.Value != nil {
				ag.CodeGen.insertData(symbol.Name, ".dword", symbol.Value)
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
