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
	cg := ag.CodeGen
	
	for _, symbol := range cg.SymTable.Symbols {
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
		} /* else { // Uninitialized variable => check expression
			var initExpr ast.Expression = varDeclr.Initializer
			var resultRegister string = ag.CodeGen.ExpressionGen.GenerateExpression(initExpr)
			fmt.Print("resultRegister: ", resultRegister, "\n") // Temporary debug print
		} */
	}

	if cg.CurrentFunction != "" && varDeclr.Initializer != nil {
		resultRegister := cg.ExpressionGen.GenerateExpression(varDeclr.Initializer)

		if offset, exists := cg.VarStackOffset[varDeclr.Name]; exists {
			if isFloatType(varDeclr.Type) {
				cg.emit("	fs %s, %d(sp)", resultRegister, offset)
			} else {
				cg.emit("	sd %s, %d(sp)", resultRegister, offset)
			}
		}

		if resultRegister != "a0" && resultRegister != "fa0" {
			cg.Registers.ReleaseRegister(resultRegister)
		}
	}
}

func isFloatType(t ast.Type) bool {
    if primitiveType, ok := t.(*ast.PrimitiveType); ok {
        return primitiveType.Name == "float"
    }
    return false
}
