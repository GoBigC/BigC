package codegen

import (
	"BigCooker/pkg/syntax/ast"
	"fmt"
)

type BranchingGenerator struct {
	CodeGen      *CodeGenerator
	labelCounter int
}

func NewBranchingGenerator(cg *CodeGenerator) *BranchingGenerator {
	return &BranchingGenerator{
		CodeGen: cg,
	}
}

func (bg *BranchingGenerator) GenerateIfStatement(stmt *ast.IfStatement) {
	cg := bg.CodeGen

	cg.emitComment("=== Begin if statement ===")
	cg.emitComment("Condition:")

	if stmt.Condition != nil {

		var condReg string

		switch cond := stmt.Condition.(type) {
		case *ast.BinaryExpression, *ast.UnaryExpression:
			condReg, _ = cg.ExpressionGen.GenerateExpression(cond)
		default:
			panic("Unsupported condition: only binary or unary expressions are allowed")
		}

		elseLabel := bg.NewLabel()
		endLabel := bg.NewLabel()

		cg.emit("beqz %s, %s", condReg, elseLabel)

		// Generate code for the then block
		cg.emitComment("Then block:")
		if stmt.ThenBlock != nil {
			cg.BlockGen.GenerateBlock(stmt.ThenBlock)
		}

		// Handle else block
		if stmt.ElseBlock != nil {
			cg.emit("j %s", endLabel) // Jump to the end after the else block
			cg.emit("%s:", elseLabel) // Else label
			switch elseBlock := stmt.ElseBlock.(type) {
			case *ast.Block:
				cg.BlockGen.GenerateBlock(elseBlock)
			case *ast.IfStatement:
				bg.GenerateIfStatement(elseBlock)
			default:
				cg.emitComment("Unsupported else block type")
			}
		} else {
			cg.emit("%s:", elseLabel)
		}

		cg.emit("%s:", endLabel)

		// Release the register used for condition
		if stmt.Condition != nil {
			cg.Registers.ReleaseRegister(condReg)
		}

		cg.emitComment("End if statement")
	} else {
		panic("Condition is null!!!")
	}
}

func (bg *BranchingGenerator) NewLabel() string {
	label := fmt.Sprintf("L%d", bg.labelCounter)
	bg.labelCounter++
	return label
}
