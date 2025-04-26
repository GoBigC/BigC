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

func (ag *AssignmentGenerator) GenerateVarDeclaration(varDeclr ast.VarDeclaration) {

	for _, symbol := range ag.CodeGen.SymTable.Symbols {
		if symbol.ReturnType != nil || symbol.Parameters != nil {
			continue // Skip function symbols
		}
		if symbol.Value != nil { // Initialized variable
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
		} else { // Uninitialized variable => check expression
			var initExpr ast.Expression = varDeclr.Initializer
			var resultRegister string = ag.CodeGen.ExpressionGen.GenerateExpression(initExpr)
			fmt.Print("resultRegister: ", resultRegister, "\n") // Temporary debug print
		}
	}
}
