package codegen

import (
	"BigCooker/pkg/syntax/ast"
	"fmt"
)

type BlockGenerator struct {
	CodeGen *CodeGenerator
}

func NewBlockGenerator(cg *CodeGenerator) *BlockGenerator {
	return &BlockGenerator{
		CodeGen: cg,
	}
}

func (bg *BlockGenerator) GenerateBlock(block *ast.Block) {
	if block == nil || len(block.Items) == 0 {
		bg.CodeGen.emitComment("Empty block")
		return
	}

	for i, item := range block.Items {
		bg.CodeGen.emitComment("Statement #%d", i+1)
		bg.GenerateBlockItem(item)
	}
}

func (bg *BlockGenerator) GenerateBlockItem(item ast.BlockItem) {
	switch stmt := item.(type) {
	case *ast.VarDeclaration:
		bg.CodeGen.AssignmentGen.GenerateVarDeclaration(*stmt)
	case *ast.ExpressionStatement:
		bg.CodeGen.ExpressionGen.GenerateExpression(stmt.Expr)
	case *ast.IfStatement:
		bg.CodeGen.BranchingGen.GenerateIfStatement(stmt)
	// case *ast.WhileStatement:
	// 	bg.CodeGen.LoopingGen.GenerateWhileStatement(stmt)
	case *ast.ReturnStatement:
		bg.CodeGen.emitComment("Return statement")
	case *ast.Block:
		bg.GenerateBlock(stmt)
	default:
		panic(fmt.Sprintf("Unsupported block item: %T", stmt))
	}
}
