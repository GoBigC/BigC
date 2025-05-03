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

func (ag *AssignmentGenerator) GenerateVarDeclaration(varDecl *ast.VarDeclaration) {
	cg := ag.CodeGen
	rp := cg.Registers

	// Use plain name (no function prefix since only main exists)
	name := varDecl.Name
	symbol, found := cg.SymTable.Lookup("main." + name)
	if !found {
		symbol, _ = cg.SymTable.Lookup(name)
	}

	// Check if local (inside main) or global
	isLocal := symbol.Scope.ValidFirstLine > 1 // Locals declared inside main
	isFloat := isFloatType(varDecl.Type)
	size := 8 // 8 bytes for int, float, bool, char
	var offset int

	if isLocal {
		// Allocate stack space
		offset = cg.AllocateStack(name, size)
	} else {
		// Global: Allocate in .data
		if isFloat {
			cg.insertData(name, ".double", 0.0)
		} else {
			cg.insertData(name, ".dword", 0)
		}
	}

	if varDecl.Initializer != nil {
		// Generate initializer
		resultReg, _ := cg.ExpressionGen.GenerateExpression(varDecl.Initializer)

		if isLocal {
			// Store to stack
			if isFloat {
				cg.emit("    fsd %s, %d(sp)", resultReg, offset)
				cg.Registers.ReleaseRegister(resultReg)
			} else {
				cg.emit("    sd %s, %d(sp)", resultReg, offset)
				cg.Registers.ReleaseRegister(resultReg)
			}
		} else {
			// Store to global
			addressReg := rp.GetTmpRegister()
			cg.emit("    la %s, %s", addressReg, name)
			if isFloat {
				cg.emit("    fsd %s, 0(%s)", resultReg, addressReg)
				cg.Registers.ReleaseRegister(resultReg)
			} else {
				cg.emit("    sd %s, 0(%s)", resultReg, addressReg)
				cg.Registers.ReleaseRegister(resultReg)
			}
			cg.Registers.ReleaseRegister(addressReg)
		}
	} else if isLocal {
		// Initialize to 0 for locals without initializer
		if isFloat {
			resultReg := rp.GetFloatTmpRegister()
			cg.emit("    fmv.d.x %s, x0", resultReg)
			cg.emit("    fsd %s, %d(sp)", resultReg, offset)
			cg.Registers.ReleaseRegister(resultReg)
		} else {
			cg.emit("    sd x0, %d(sp)", offset)
		}
	}

	fmt.Printf("Var %s: local=%v, offset=%d\n", name, isLocal, offset)
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

	// Generate value expression
	resultReg, _ := eg.GenerateExpression(value)

	// Look up symbol
	name := id.Name
	symbol, found := cg.SymTable.Lookup("main." + name)
	if !found {
		symbol, _ = cg.SymTable.Lookup(name)
	}

	// Determine if local or global
	isLocal := symbol.Scope.ValidFirstLine > 1 // Locals declared inside main
	isFloat := isFloatType(symbol.Type)

	if isLocal {
		// Local: Store to stack offset
		offset := cg.GetStackOffset(name)
		if isFloat {
			cg.emit("    fsd %s, %d(sp)", resultReg, offset)
			cg.Registers.ReleaseRegister(resultReg)
		} else {
			cg.emit("    sd %s, %d(sp)", resultReg, offset)
			cg.Registers.ReleaseRegister(resultReg)
		}
	} else {
		// Global: Store to .data
		addressReg := rp.GetTmpRegister()
		cg.emit("    la %s, %s", addressReg, name)
		if isFloat {
			cg.emit("    fsd %s, 0(%s)", resultReg, addressReg)
			cg.Registers.ReleaseRegister(resultReg)
		} else {
			cg.emit("    sd %s, 0(%s)", resultReg, addressReg)
			cg.Registers.ReleaseRegister(resultReg)
		}
		cg.Registers.ReleaseRegister(addressReg)
	}
}
