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
				fmt.Printf("variable %s initializer is %s\n", varDecl.Name, varDecl.Initializer)
				// First, add the variable to .data section
				if isFloatType(varDecl.Type) {
					fmt.Printf("inserting float\n")
					cg.insertData(varDecl.Name, ".double", 0.0)
				} else {
					fmt.Printf("inserting others: %s\n", varDecl.Name)
					cg.insertData(varDecl.Name, ".dword", 0)
				}

				resultRegister, _ := cg.ExpressionGen.GenerateExpression(varDecl.Initializer)
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

func (ag *AssignmentGenerator) GenerateArrayAssignment(arrExpr *ast.ArrayAccessExpression, value ast.Expression) {
	cg := ag.CodeGen
	eg := cg.ExpressionGen
    rp := cg.Registers

	valueRegister, _ := eg.GenerateExpression(value)
	elemAddrRegister, indexRegister, elemType := eg.CalculateArrayElementAddress(arrExpr.Array, arrExpr.Index)

	isFloat := false
    if primType, ok := elemType.(*ast.PrimitiveType); ok {
        isFloat = primType.Name == "float"
    }

	if isFloat {
		cg.emit("	fsd %s, 0(%s)", valueRegister, elemAddrRegister)
	} else {
		cg.emit("	sd %s, 0(%s)", valueRegister, elemAddrRegister)
	}

	if elemAddrRegister != "a0" && elemAddrRegister != "fa0" {
		rp.ReleaseRegister(elemAddrRegister)
	}
    if indexRegister != "a0" && indexRegister != "fa0" {
        rp.ReleaseRegister(indexRegister)
    }
    if valueRegister != "a0" && valueRegister != "fa0" {
        rp.ReleaseRegister(valueRegister)
    }
}

func (ag *AssignmentGenerator) GenerateVariableAssignment(id *ast.Identifier, value ast.Expression) {
    cg := ag.CodeGen
    eg := cg.ExpressionGen
    rp := cg.Registers
    
    resultReg, resultType := eg.GenerateExpression(value)
    
    addressReg := rp.GetTmpRegister()
    cg.emit("    la %s, %s", addressReg, id.Name)
    
    if primType, ok := resultType.(*ast.PrimitiveType); ok && primType.Name == "float" {
        cg.emit("    fsd %s, 0(%s)", resultReg, addressReg)
    } else {
        cg.emit("    sd %s, 0(%s)", resultReg, addressReg)
    }
    
    rp.ReleaseRegister(addressReg)
    if resultReg != "a0" && resultReg != "fa0" {
        rp.ReleaseRegister(resultReg)
    }
}