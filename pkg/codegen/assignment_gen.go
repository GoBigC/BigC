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

func (ag *AssignmentGenerator) GenerateVarDeclaration(varDecl ast.VarDeclaration) {
	cg := ag.CodeGen

	switch t := varDecl.Type.(type) {
	case *ast.PrimitiveType: 
		if varDecl.Initializer != nil {
			if lit, ok := varDecl.Initializer.(*ast.IntegerLiteral); ok {
				// int literal
				cg.insertData(varDecl.Name, ".dword", lit.Value)
			} else if lit, ok := varDecl.Initializer.(*ast.FloatLiteral); ok {
				// float literal 
				cg.insertData(varDecl.Name, ".double", lit.Value)
			} else if lit, ok := varDecl.Initializer.(*ast.BoolLiteral); ok {
				// bool literal --> int
				value := 0 
				if lit.Value {
					value = 1
				}
				cg.insertData(varDecl.Name, ".dword", value)
			} else if lit, ok := varDecl.Initializer.(*ast.CharLiteral); ok {
				// char literal --> int ascii code 
				charAsciiCode := getAscii(lit)
				cg.insertData(varDecl.Name, ".dword", charAsciiCode)
			} else { // expression
				resultRegister := cg.ExpressionGen.GenerateExpression(varDecl.Initializer)
				addressRegister := cg.Registers.GetTmpRegister()

				cg.emit("	la %s, %s", addressRegister, varDecl.Name)
				
				if isFloatType(varDecl.Type) { 
					cg.emit("	fsd %s, 0(%s)", resultRegister, addressRegister)
				} else {
					cg.emit("	sd %s, 0(%s)", resultRegister, addressRegister)
				}

				cg.Registers.ReleaseRegister(addressRegister)
				if resultRegister != "a0" && resultRegister != "fa0" {
                    cg.Registers.ReleaseRegister(resultRegister)
                }

				if isFloatType(varDecl.Type) {
                    cg.insertData(varDecl.Name, ".double", 0.0)
                } else {
                    cg.insertData(varDecl.Name, ".dword", 0)
                }
			}
		} else {
            if t.Name == "float" {
                ag.CodeGen.insertData(varDecl.Name, ".double", 0.0)
            } else {
                ag.CodeGen.insertData(varDecl.Name, ".dword", 0)
            }
        }
	case *ast.ArrayType:
		var id string = varDecl.Name 
		fmt.Printf("------ id is %s", id)
		symbol, found := cg.SymTable.Lookup("main." + id) // try local 
		if !found {
			symbol, found = cg.SymTable.Lookup(id) // try global
		}

		if found && symbol.ArraySize > 0{
			cg.insertData(varDecl.Name, ".space", 8*symbol.ArraySize)
		} else {
			panic(fmt.Sprintf("Invalid array size at line %d", varDecl.Line))
		}
	default: 
		panic(fmt.Sprintf("Unsupported type for variable declaration: %T", t))
	}
}

func getAscii(lit *ast.CharLiteral) int {
	return int(lit.Value)
}

func isFloatType(typeExpr ast.Type) bool {
	if primitiveType, ok := typeExpr.(*ast.PrimitiveType); ok {
		return primitiveType.Name == "float"
	}
	return false
}
