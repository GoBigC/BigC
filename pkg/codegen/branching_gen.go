package codegen

import (
	"BigCooker/pkg/syntax/ast" // Import your symbol table package
)

type BranchingGenerator struct {
	CodeGen *CodeGenerator
}

func NewBranchingGenerator(cg *CodeGenerator) *BranchingGenerator {
	return &BranchingGenerator{
		CodeGen: cg,
	}
}

func (bg *BranchingGenerator) GenerateIfStatement(stmt *ast.IfStatement) {
	cg := bg.CodeGen

	cg.emitComment("=== Begin if statement ===")

	var condReg = "t1"

	elseLabel := cg.NewLabel()
	endLabel := cg.NewLabel()

	cg.emit("beq %s, x0, %s", condReg, elseLabel)

	// Generate code for the then block
	cg.emitComment("Then block:")
	if stmt.ThenBlock != nil {
		bg.GenerateBlock(stmt.ThenBlock)
	}

	// Handle else block
	if stmt.ElseBlock != nil {
		cg.emit("j %s", endLabel) // Jump to the end after the else block
		cg.emit("%s:", elseLabel) // Else label
		if elseBlock, ok := stmt.ElseBlock.(*ast.Block); ok {
			bg.GenerateBlock(elseBlock)
		} else {
			cg.emitComment("Else block is null")
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
}

func (bg *BranchingGenerator) GenerateBlock(block *ast.Block) {
	for i, item := range block.Items {
		bg.CodeGen.emitComment("Statement #%d", i+1)
		bg.GenerateBlockItem(item)
	}
}

func (bg *BranchingGenerator) GenerateBlockItem(item ast.BlockItem) {
	switch stmt := item.(type) {
	// case *ast.ExpressionStatement:
	// 	bg.GenerateExpressionStatement(stmt)
	case *ast.VarDeclaration:
		bg.CodeGen.AssignmentGen.GenerateVarDeclaration(*stmt)
	case *ast.IfStatement:
		bg.GenerateIfStatement(stmt)
	// Add more statement types as needed
	default:
		bg.CodeGen.emitComment("Unsupported statement type: %T", item)
	}
}
